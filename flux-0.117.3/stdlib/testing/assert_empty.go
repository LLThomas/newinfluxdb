package testing

import (
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/execute"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/plan"
	"github.com/influxdata/flux/runtime"
	"sync"
)

const AssertEmptyKind = "assertEmpty"

type AssertEmptyOpSpec struct{}

func (s *AssertEmptyOpSpec) Kind() flux.OperationKind {
	return AssertEmptyKind
}

func init() {
	assertEmptySignature := runtime.MustLookupBuiltinType("testing", "assertEmpty")

	runtime.RegisterPackageValue("testing", "assertEmpty", flux.MustValue(flux.FunctionValue(AssertEmptyKind, createAssertEmptyOpSpec, assertEmptySignature)))
	flux.RegisterOpSpec(AssertEmptyKind, newAssertEmptyOp)
	plan.RegisterProcedureSpec(AssertEmptyKind, newAssertEmptyProcedure, AssertEmptyKind)
	execute.RegisterTransformation(AssertEmptyKind, createAssertEmptyTransformation)
}

func createAssertEmptyOpSpec(args flux.Arguments, a *flux.Administration) (flux.OperationSpec, error) {
	if err := a.AddParentFromArgs(args); err != nil {
		return nil, err
	}
	return &AssertEmptyOpSpec{}, nil
}

func newAssertEmptyOp() flux.OperationSpec {
	return new(AssertEmptyOpSpec)
}

type AssertEmptyProcedureSpec struct {
	plan.DefaultCost
}

func (s *AssertEmptyProcedureSpec) Kind() plan.ProcedureKind {
	return AssertEmptyKind
}

func (s *AssertEmptyProcedureSpec) Copy() plan.ProcedureSpec {
	ns := *s
	return &ns
}

func newAssertEmptyProcedure(qs flux.OperationSpec, pa plan.Administration) (plan.ProcedureSpec, error) {
	if _, ok := qs.(*AssertEmptyOpSpec); !ok {
		return nil, errors.Newf(codes.Internal, "invalid spec type %T", qs)
	}
	return &AssertEmptyProcedureSpec{}, nil
}

type AssertEmptyTransformation struct {
	execute.ExecutionNode
	failures int64

	d     execute.Dataset
	cache execute.TableBuilderCache
}

func (t *AssertEmptyTransformation) SetRoad(m map[string]int, m2 map[string]string, transformation *execute.Transformation, state *execute.ExecutionState) {
	panic("implement me")
}

func (t *AssertEmptyTransformation) GetRoad(s string, i int) (*execute.ConsecutiveTransport, *execute.Transformation) {
	panic("implement me")
}

func (t *AssertEmptyTransformation) GetEs() *execute.ExecutionState {
	panic("implement me")
}

func (t *AssertEmptyTransformation) SetWG(WG *sync.WaitGroup) {
	panic("implement me")
}

func (t *AssertEmptyTransformation) ProcessTbl(id execute.DatasetID, tbls []flux.Table) error {
	panic("implement me")
}

func (t *AssertEmptyTransformation) ClearCache() error {
	panic("implement me")
}

func createAssertEmptyTransformation(id execute.DatasetID, mode execute.AccumulationMode, spec plan.ProcedureSpec, a execute.Administration, whichPipeThread int) (execute.Transformation, execute.Dataset, error) {
	cache := execute.NewTableBuilderCache(a.Allocator())
	dataset := execute.NewDataset(id, mode, cache)
	if _, ok := spec.(*AssertEmptyProcedureSpec); !ok {
		return nil, nil, errors.Newf(codes.Internal, "invalid spec type %T", spec)
	}

	transform := NewAssertEmptyTransformation(dataset, cache)
	return transform, dataset, nil
}

func NewAssertEmptyTransformation(d execute.Dataset, cache execute.TableBuilderCache) *AssertEmptyTransformation {
	return &AssertEmptyTransformation{
		d:     d,
		cache: cache,
	}
}

func (t *AssertEmptyTransformation) RetractTable(id execute.DatasetID, key flux.GroupKey) error {
	return t.d.RetractTable(key)
}

func (t *AssertEmptyTransformation) Process(id execute.DatasetID, tbl flux.Table) error {
	if !tbl.Empty() {
		t.failures++
	}
	// TODO: The Do method must be called at the moment.
	return tbl.Do(func(cr flux.ColReader) error {
		return nil
	})
}

func (t *AssertEmptyTransformation) UpdateWatermark(id execute.DatasetID, mark execute.Time) error {
	return t.d.UpdateWatermark(mark)
}

func (t *AssertEmptyTransformation) UpdateProcessingTime(id execute.DatasetID, mark execute.Time) error {
	return t.d.UpdateProcessingTime(mark)
}

func (t *AssertEmptyTransformation) Finish(id execute.DatasetID, err error, windowModel bool) {
	if err == nil && t.failures > 0 {
		err = errors.Newf(codes.Aborted, "found %d tables that were not empty", t.failures)
	}
	t.d.Finish(err, windowModel)
}
