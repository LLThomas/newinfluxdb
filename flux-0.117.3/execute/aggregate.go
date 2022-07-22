package execute

import (
	"github.com/apache/arrow/go/arrow/array"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/plan"
	"log"
	"os"
	"sync"
)

type aggregateTransformation struct {
	ExecutionNode
	d     Dataset
	cache TableBuilderCache
	agg   Aggregate

	config AggregateConfig

	whichPipeThread int
	WG *sync.WaitGroup
	OperatorIndex 	map[string]int
	OperatorMap 	map[string]string
	ResOperator		*Transformation
	ExecutionState 	*ExecutionState
}

func (t *aggregateTransformation) SetRoad(m map[string]int, m2 map[string]string, transformation *Transformation, state *ExecutionState) {
	t.OperatorIndex = m
	t.OperatorMap = m2
	t.ResOperator = transformation
	t.ExecutionState = state
}

func (t *aggregateTransformation) GetRoad(label string, thread int) (*ConsecutiveTransport, *Transformation) {
	nextOperatorString := ""
	if t.OperatorMap[label] != "" {
		nextOperatorString = t.OperatorMap[label]
	}
	if nextOperatorString == "" {
		return nil, t.ResOperator
	}
	return t.ExecutionState.ESmultiThreadPipeLine[thread].Worker[t.OperatorIndex[nextOperatorString]], t.ResOperator
}

func (t *aggregateTransformation) GetEs() *ExecutionState {
	return t.ExecutionState
}

func (t *aggregateTransformation) SetWG(WG *sync.WaitGroup) {
	panic("implement me")
}

func (t *aggregateTransformation) ClearCache() error {
	return t.d.ClearCache()
}

type AggregateConfig struct {
	plan.DefaultCost
	Columns []string `json:"columns"`
}

var DefaultAggregateConfig = AggregateConfig{
	Columns: []string{DefaultValueColLabel},
}

func (c AggregateConfig) Copy() AggregateConfig {
	nc := c
	if c.Columns != nil {
		nc.Columns = make([]string, len(c.Columns))
		copy(nc.Columns, c.Columns)
	}
	return nc
}

func (c *AggregateConfig) ReadArgs(args flux.Arguments) error {
	if col, ok, err := args.GetString("column"); err != nil {
		return err
	} else if ok {
		c.Columns = []string{col}
	} else {
		c.Columns = DefaultAggregateConfig.Columns
	}
	return nil
}

func NewAggregateTransformation(d Dataset, c TableBuilderCache, agg Aggregate, config AggregateConfig, whichPipeThread int) *aggregateTransformation {
	return &aggregateTransformation{
		d:      d,
		cache:  c,
		agg:    agg,
		config: config,
		whichPipeThread: whichPipeThread,
	}
}

func NewAggregateTransformationAndDataset(id DatasetID, mode AccumulationMode, agg Aggregate, config AggregateConfig, a *memory.Allocator, whichPipeThread int) (*aggregateTransformation, Dataset) {
	cache := NewTableBuilderCache(a)
	d := NewDataset(id, mode, cache)
	return NewAggregateTransformation(d, cache, agg, config, whichPipeThread), d
}

func (t *aggregateTransformation) RetractTable(id DatasetID, key flux.GroupKey) error {
	//TODO(nathanielc): Store intermediate state for retractions
	return t.d.RetractTable(key)
}

func (t *aggregateTransformation) ProcessTbl(id DatasetID, tbls []flux.Table) error {

	//log.Println("aggregate: ")
	//for k := 0; k < len(tbls); k++ {
	//	log.Println(tbls[k].Key())
	//}

	//nextOperator := FindNextOperator(t.Label(), t.whichPipeThread)
	//resOperator := ResOperator
	nextOperator, resOperator := t.GetRoad(t.Label(), t.whichPipeThread)

	var tables []flux.Table
	for i := 0; i < len(tbls); i++ {
		tbl := tbls[i]

		builder, created := t.cache.TableBuilder(tbl.Key())
		if !created {
			return errors.Newf(codes.FailedPrecondition, "aggregate found duplicate table with key: %v", tbl.Key())
		}

		if err := AddTableKeyCols(tbl.Key(), builder); err != nil {
			return err
		}

		builderColMap := make([]int, len(t.config.Columns))
		tableColMap := make([]int, len(t.config.Columns))
		aggregates := make([]ValueFunc, len(t.config.Columns))
		cols := tbl.Cols()

		should_continue := false
		for j, label := range t.config.Columns {
			idx := -1
			for bj, bc := range cols {
				if bc.Label == label {
					idx = bj
					break
				}
			}
			if idx < 0 {
				should_continue = true
				break;
				//return errors.Newf(codes.FailedPrecondition, "column %q does not exist", label)
			}
			c := cols[idx]
			if tbl.Key().HasCol(c.Label) {
				return errors.New(codes.FailedPrecondition, "cannot aggregate columns that are part of the group key")
			}
			var vf ValueFunc
			switch c.Type {
			case flux.TBool:
				vf = t.agg.NewBoolAgg()
			case flux.TInt:
				vf = t.agg.NewIntAgg()
			case flux.TUInt:
				vf = t.agg.NewUIntAgg()
			case flux.TFloat:
				vf = t.agg.NewFloatAgg()
			case flux.TString:
				vf = t.agg.NewStringAgg()
			}
			if vf == nil {
				return errors.Newf(codes.FailedPrecondition, "unsupported aggregate column type %v", c.Type)
			}
			aggregates[j] = vf

			var err error
			builderColMap[j], err = builder.AddCol(flux.ColMeta{
				Label: c.Label,
				Type:  vf.Type(),
			})
			if err != nil {
				return err
			}
			tableColMap[j] = idx
		}
		if should_continue {
			continue
		}

		// just return t when calling BlockIterator for the first time
		cr, _ := tbl.BlockIterator(0)

		// If cr is nil, it means we loss the data in original data file.
		// In case of interrupt the program, we neglect this condition.
		if cr == nil {
			tbl.BlockIterator(1)
			continue
		}

		// use cr
		for j := range t.config.Columns {
			vf := aggregates[j]

			tj := tableColMap[j]
			c := tbl.Cols()[tj]

			switch c.Type {
			case flux.TBool:
				vf.(DoBoolAgg).DoBool(cr.Bools(tj))
			case flux.TInt:
				vf.(DoIntAgg).DoInt(cr.Ints(tj))
			case flux.TUInt:
				vf.(DoUIntAgg).DoUInt(cr.UInts(tj))
			case flux.TFloat:

				if cr == nil {
					log.Println("nil: ", t.label)
					os.Exit(0)
				}

				vf.(DoFloatAgg).DoFloat(cr.Floats(tj))
			case flux.TString:
				vf.(DoStringAgg).DoString(cr.Strings(tj))
			default:
				return errors.Newf(codes.Invalid, "unsupported aggregate type %v", c.Type)
			}
		}
		for j, vf := range aggregates {
			bj := builderColMap[j]

			// If the value is null, append a null to the column.
			if vf.IsNull() {
				if err := builder.AppendNil(bj); err != nil {
					return err
				}
				continue
			}

			// Append aggregated value
			switch vf.Type() {
			case flux.TBool:
				v := vf.(BoolValueFunc).ValueBool()
				if err := builder.AppendBool(bj, v); err != nil {
					return err
				}
			case flux.TInt:
				v := vf.(IntValueFunc).ValueInt()
				if err := builder.AppendInt(bj, v); err != nil {
					return err
				}
			case flux.TUInt:
				v := vf.(UIntValueFunc).ValueUInt()
				if err := builder.AppendUInt(bj, v); err != nil {
					return err
				}
			case flux.TFloat:
				v := vf.(FloatValueFunc).ValueFloat()
				if err := builder.AppendFloat(bj, v); err != nil {
					return err
				}
			case flux.TString:
				v := vf.(StringValueFunc).ValueString()
				if err := builder.AppendString(bj, v); err != nil {
					return err
				}
			}
		}

		if err := AppendKeyValues(tbl.Key(), builder); err != nil {
			return err
		}

		table, _ := builder.Table()
		tables = append(tables, table)

		// release the table memory when calling BlockIterator next time
		tbl.BlockIterator(1)
		t.ClearCache()
	}

	// send table to next operator
	// If tables is nil, error (pipe worker error:  close of closed channel) will occur because of multiple close operations of close(t.finished)
	// in transport.go.
	if tables != nil && len(tables) > 0 {
		if nextOperator == nil {
			(*resOperator).ProcessTbl(DatasetID{0}, tables)
		} else {
			nextOperator.PushToChannel(tables)
		}
	}

	return nil
}

func (t *aggregateTransformation) Process(id DatasetID, tbl flux.Table) error {

	builder, created := t.cache.TableBuilder(tbl.Key())
	if !created {
		return errors.Newf(codes.FailedPrecondition, "aggregate found duplicate table with key: %v", tbl.Key())
	}

	if err := AddTableKeyCols(tbl.Key(), builder); err != nil {
		return err
	}

	builderColMap := make([]int, len(t.config.Columns))
	tableColMap := make([]int, len(t.config.Columns))
	aggregates := make([]ValueFunc, len(t.config.Columns))
	cols := tbl.Cols()
	for j, label := range t.config.Columns {
		idx := -1
		for bj, bc := range cols {
			if bc.Label == label {
				idx = bj
				break
			}
		}
		if idx < 0 {
			return errors.Newf(codes.FailedPrecondition, "column %q does not exist", label)
		}
		c := cols[idx]
		if tbl.Key().HasCol(c.Label) {
			return errors.New(codes.FailedPrecondition, "cannot aggregate columns that are part of the group key")
		}
		var vf ValueFunc
		switch c.Type {
		case flux.TBool:
			vf = t.agg.NewBoolAgg()
		case flux.TInt:
			vf = t.agg.NewIntAgg()
		case flux.TUInt:
			vf = t.agg.NewUIntAgg()
		case flux.TFloat:
			vf = t.agg.NewFloatAgg()
		case flux.TString:
			vf = t.agg.NewStringAgg()
		}
		if vf == nil {
			return errors.Newf(codes.FailedPrecondition, "unsupported aggregate column type %v", c.Type)
		}
		aggregates[j] = vf

		var err error
		builderColMap[j], err = builder.AddCol(flux.ColMeta{
			Label: c.Label,
			Type:  vf.Type(),
		})
		if err != nil {
			return err
		}
		tableColMap[j] = idx
	}

	if err := tbl.Do(func(cr flux.ColReader) error {
		for j := range t.config.Columns {
			vf := aggregates[j]

			tj := tableColMap[j]
			c := tbl.Cols()[tj]

			switch c.Type {
			case flux.TBool:
				vf.(DoBoolAgg).DoBool(cr.Bools(tj))
			case flux.TInt:
				vf.(DoIntAgg).DoInt(cr.Ints(tj))
			case flux.TUInt:
				vf.(DoUIntAgg).DoUInt(cr.UInts(tj))
			case flux.TFloat:
				vf.(DoFloatAgg).DoFloat(cr.Floats(tj))
			case flux.TString:
				vf.(DoStringAgg).DoString(cr.Strings(tj))
			default:
				return errors.Newf(codes.Invalid, "unsupported aggregate type %v", c.Type)
			}
		}
		return nil
	}); err != nil {
		return err
	}
	for j, vf := range aggregates {
		bj := builderColMap[j]

		// If the value is null, append a null to the column.
		if vf.IsNull() {
			if err := builder.AppendNil(bj); err != nil {
				return err
			}
			continue
		}

		// Append aggregated value
		switch vf.Type() {
		case flux.TBool:
			v := vf.(BoolValueFunc).ValueBool()
			if err := builder.AppendBool(bj, v); err != nil {
				return err
			}
		case flux.TInt:
			v := vf.(IntValueFunc).ValueInt()
			if err := builder.AppendInt(bj, v); err != nil {
				return err
			}
		case flux.TUInt:
			v := vf.(UIntValueFunc).ValueUInt()
			if err := builder.AppendUInt(bj, v); err != nil {
				return err
			}
		case flux.TFloat:
			v := vf.(FloatValueFunc).ValueFloat()
			if err := builder.AppendFloat(bj, v); err != nil {
				return err
			}
		case flux.TString:
			v := vf.(StringValueFunc).ValueString()
			if err := builder.AppendString(bj, v); err != nil {
				return err
			}
		}
	}

	err := AppendKeyValues(tbl.Key(), builder)
	//b, _ := builder.Table()
	//nextOperator := OperatorMap[t.Label()]
	//if nextOperator == nil {
	//	ResOperator.Process(DatasetID{0}, b)
	//} else {
	//	//nextOperator.PushToChannel(b)
	//}
	return err
}

func (t *aggregateTransformation) UpdateWatermark(id DatasetID, mark Time) error {
	return t.d.UpdateWatermark(mark)
}
func (t *aggregateTransformation) UpdateProcessingTime(id DatasetID, pt Time) error {
	return t.d.UpdateProcessingTime(pt)
}
func (t *aggregateTransformation) Finish(id DatasetID, err error, windowModel bool) {
	t.d.Finish(err, windowModel)
}

type Aggregate interface {
	NewBoolAgg() DoBoolAgg
	NewIntAgg() DoIntAgg
	NewUIntAgg() DoUIntAgg
	NewFloatAgg() DoFloatAgg
	NewStringAgg() DoStringAgg
}

type ValueFunc interface {
	Type() flux.ColType
	IsNull() bool
}
type DoBoolAgg interface {
	ValueFunc
	DoBool(*array.Boolean)
}
type DoFloatAgg interface {
	ValueFunc
	DoFloat(*array.Float64)
}
type DoIntAgg interface {
	ValueFunc
	DoInt(*array.Int64)
}
type DoUIntAgg interface {
	ValueFunc
	DoUInt(*array.Uint64)
}
type DoStringAgg interface {
	ValueFunc
	DoString(*array.Binary)
}

type BoolValueFunc interface {
	ValueBool() bool
}
type FloatValueFunc interface {
	ValueFloat() float64
}
type IntValueFunc interface {
	ValueInt() int64
}
type UIntValueFunc interface {
	ValueUInt() uint64
}
type StringValueFunc interface {
	ValueString() string
}
