// Package execute contains the implementation of the execution phase in the query engine.
package execute

import (
	"context"
	"fmt"
	"github.com/influxdata/influxdb/v2/tsdb/cursors"
	"log"
	"reflect"
	"runtime/debug"
	"sync"
	"time"

	"github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/internal/errors"
	"github.com/influxdata/flux/memory"
	"github.com/influxdata/flux/metadata"
	"github.com/influxdata/flux/plan"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type Executor interface {
	// Execute will begin execution of the plan.Spec using the memory allocator.
	// This returns a mapping of names to the query results.
	// This will also return a channel for the Metadata from the query. The channel
	// may return zero or more values. The returned channel must not require itself to
	// be read so the executor must allocate enough space in the channel so if the channel
	// is unread that it will not block.
	Execute(ctx context.Context, p *plan.Spec, a *memory.Allocator) (map[string]flux.Result, <-chan metadata.Metadata, error)
}

type executor struct {
	logger *zap.Logger
}

func NewExecutor(logger *zap.Logger) Executor {
	if logger == nil {
		logger = zap.NewNop()
	}
	e := &executor{
		logger: logger,
	}
	return e
}

type streamContext struct {
	bounds *Bounds
}

func (ctx streamContext) Bounds() *Bounds {
	return ctx.bounds
}

type executionState struct {
	p *plan.Spec

	ctx    context.Context
	cancel func()
	alloc  *memory.Allocator

	resources flux.ResourceManagement

	results map[string]flux.Result
	sources []Source
	metaCh  chan metadata.Metadata

	transports []Transport

	dispatcher *poolDispatcher
	logger     *zap.Logger

	consecutiveTransportSet []*consecutiveTransport

	// set the number of pipeline to 6 which equals to the number of cores
	// each pipeline has a set of pipe workers
	ESmultiThreadPipeLine []*MultiThreadPipeLine

	// number of finish operator
	// when the sum is equal to the number of thread, result operator should be shut down
	// in case of sending duplicate finish msgs to finish operator
	numFinishMsgCount int32

	// size of a block group
	Len int
}

func (e *executor) Execute(ctx context.Context, p *plan.Spec, a *memory.Allocator) (map[string]flux.Result, <-chan metadata.Metadata, error) {
	es, err := e.createExecutionState(ctx, p, a)
	if err != nil {
		return nil, nil, errors.Wrap(err, codes.Inherit, "failed to initialize execute state")
	}

	log.Println("number of pipeline: ", len(es.ESmultiThreadPipeLine))
	log.Println("number of operator in each pipeline: ", len(es.ESmultiThreadPipeLine[0].Worker))
	// construct global operators line
	n := len(es.consecutiveTransportSet)
	for i := 0; i < n; i++ {
		e := es.consecutiveTransportSet[i]
		OperatorIndex[e.Label()] = i
		log.Println("OperatorIndex: ", i, e.Label())
		if i+1 < n {
			OperatorMap[e.Label()] = es.consecutiveTransportSet[i+1].Label()
		} else {
			OperatorMap[e.Label()] = ""
		}
		log.Println("OperatorMap: ", i, OperatorMap[e.Label()])
	}

	//for i := 0; i < len(es.ESmultiThreadPipeLine); i++ {
	//	for j := 0; j < len(es.ESmultiThreadPipeLine[i].Worker); j++ {
	//		log.Printf("Execute: %s %p ", es.ESmultiThreadPipeLine[i].Worker[j].t.Label(), es.ESmultiThreadPipeLine[i].Worker[j].t)
	//	}
	//}

	// start pipe worker in each pipeline
	for i := 0; i < len(es.ESmultiThreadPipeLine); i++ {
		es.ESmultiThreadPipeLine[i].startPipeLine(ctx)
	}

	// for put series key into executionState.multiPipeLine, set es to global ExecutionState in multi_thread_pipe.go
	// TODO: this is a bad idea, find elegant approach to solve it
	ExecutionState = es

	es.do()
	return es.results, es.metaCh, nil
}

func validatePlan(p *plan.Spec) error {
	if p.Resources.ConcurrencyQuota == 0 {
		return errors.New(codes.Invalid, "plan must have a non-zero concurrency quota")
	}
	return nil
}

func (e *executor) createExecutionState(ctx context.Context, p *plan.Spec, a *memory.Allocator) (*executionState, error) {
	if err := validatePlan(p); err != nil {
		return nil, errors.Wrap(err, codes.Invalid, "invalid plan")
	}

	ctx, cancel := context.WithCancel(ctx)
	es := &executionState{
		p:         p,
		ctx:       ctx,
		cancel:    cancel,
		alloc:     a,
		resources: p.Resources,
		results:   make(map[string]flux.Result),
		// TODO(nathanielc): Have the planner specify the dispatcher throughput
		dispatcher:            newPoolDispatcher(10, e.logger),
		logger:                e.logger,
		// assume that number of thread is equal to the number of cores
		ESmultiThreadPipeLine: make([]*MultiThreadPipeLine, 6),
		// in case of sending duplicate finishMsg to finish operator, we count the finishMsg
		// if numFinishMsgCount equals to the number of pipeline thread, we send a realy finish msg to finish operator
		numFinishMsgCount: 0,
		// assume that the size of block group is 3
		Len: 3,
	}

	for i := 0; i < len(es.ESmultiThreadPipeLine); i++ {
		es.ESmultiThreadPipeLine[i] = newMultiPipeLine(make([]cursors.Cursor, 0), make([]cursors.Cursor, 0), make([]*consecutiveTransport, 0))
	}

	v := &createExecutionNodeVisitor{
		es:    es,
		nodes: make(map[plan.Node]Node),
	}

	if err := p.BottomUpWalk(v.Visit); err != nil {
		return nil, err
	}

	// Only sources can be a MetadataNode at the moment so allocate enough
	// space for all of them to report metadata. Not all of them will necessarily
	// report metadata.
	es.metaCh = make(chan metadata.Metadata, len(es.sources))

	return v.es, nil
}

// createExecutionNodeVisitor visits each node in a physical query plan
// and creates a node responsible for executing that physical operation.
type createExecutionNodeVisitor struct {
	es    *executionState
	nodes map[plan.Node]Node
}

func skipYields(pn plan.Node) plan.Node {
	isYield := func(pn plan.Node) bool {
		_, ok := pn.ProcedureSpec().(plan.YieldProcedureSpec)
		return ok
	}

	for isYield(pn) {
		pn = pn.Predecessors()[0]
	}

	return pn
}

func nonYieldPredecessors(pn plan.Node) []plan.Node {
	nodes := make([]plan.Node, len(pn.Predecessors()))
	for i, pred := range pn.Predecessors() {
		nodes[i] = skipYields(pred)
	}

	return nodes
}

// Visit creates the node that will execute a particular plan node
func (v *createExecutionNodeVisitor) Visit(node plan.Node) error {
	ppn, ok := node.(*plan.PhysicalPlanNode)
	if !ok {
		return fmt.Errorf("cannot execute plan node of type %T", node)
	}
	spec := node.ProcedureSpec()
	kind := spec.Kind()
	id := DatasetIDFromNodeID(node.ID())

	if yieldSpec, ok := spec.(plan.YieldProcedureSpec); ok {
		r := newResult(yieldSpec.YieldName())
		v.es.results[yieldSpec.YieldName()] = r
		v.nodes[skipYields(node)].AddTransformation(r)

		ResOperator = r

		return nil
	}

	// Add explicit stream context if bounds are set on this node
	var streamContext streamContext
	if node.Bounds() != nil {
		streamContext.bounds = &Bounds{
			Start: node.Bounds().Start,
			Stop:  node.Bounds().Stop,
		}
	}

	// Build execution context
	ec := executionContext{
		es:            v.es,
		parents:       make([]DatasetID, len(node.Predecessors())),
		streamContext: streamContext,
	}

	for i, pred := range nonYieldPredecessors(node) {
		ec.parents[i] = DatasetIDFromNodeID(pred.ID())
	}

	// If node is a leaf, create a source
	if len(node.Predecessors()) == 0 {
		createSourceFn, ok := procedureToSource[kind]

		if !ok {
			return fmt.Errorf("unsupported source kind %v", kind)
		}

		source, err := createSourceFn(spec, id, ec)

		if err != nil {
			return err
		}

		source.SetLabel(string(node.ID()))
		v.es.sources = append(v.es.sources, source)
		v.nodes[node] = source
	} else {

		// If node is internal, create a transformation.
		// For each predecessor, add a transport for sending data upstream.
		createTransformationFn, ok := procedureToTransformation[kind]

		if !ok {
			return fmt.Errorf("unsupported procedure %v", kind)
		}

		tr, ds, err := createTransformationFn(id, DiscardingMode, spec, ec, 0)

		if err != nil {
			return err
		}

		if ds, ok := ds.(DatasetContext); ok {
			ds.WithContext(v.es.ctx)
		}

		tr.SetLabel(string(node.ID()))
		if ppn.TriggerSpec == nil {
			ppn.TriggerSpec = plan.DefaultTriggerSpec
		}
		ds.SetTriggerSpec(ppn.TriggerSpec)
		v.nodes[node] = ds

		for _, p := range nonYieldPredecessors(node) {
			executionNode := v.nodes[p]
			transport := newConsecutiveTransport(v.es.ctx, v.es.dispatcher, tr, node, v.es.logger)

			v.es.transports = append(v.es.transports, transport)
			executionNode.AddTransformation(transport)

			v.es.consecutiveTransportSet = append(v.es.consecutiveTransportSet, transport)
			// add transport to each pipeline
			for i := 0; i < len(v.es.ESmultiThreadPipeLine); i++ {

				newtr, newds, newerr := createTransformationFn(id, DiscardingMode, spec, ec, i)
				if newerr != nil {
					return err
				}
				if newds, ok := newds.(DatasetContext); ok {
					newds.WithContext(v.es.ctx)
				}
				newtr.SetLabel(string(node.ID()))
				if ppn.TriggerSpec == nil {
					ppn.TriggerSpec = plan.DefaultTriggerSpec
				}
				newds.SetTriggerSpec(ppn.TriggerSpec)

				t := newConsecutiveTransport(v.es.ctx, v.es.dispatcher, newtr, node, v.es.logger)
				// set whichPipeThread value
				t.whichPipeThread = i

				//log.Printf("newConsecutiveTransport: %p\n", t.t)

				v.es.ESmultiThreadPipeLine[i].Worker = append(v.es.ESmultiThreadPipeLine[i].Worker, t)
			}

		}

		if plan.HasSideEffect(spec) && len(node.Successors()) == 0 {
			name := string(node.ID())
			r := newResult(name)
			v.es.results[name] = r
			v.nodes[skipYields(node)].AddTransformation(r)
		}
	}

	return nil
}

func (es *executionState) abort(err error) {
	for _, r := range es.results {
		r.(*result).abort(err)
	}
	es.cancel()
}

func (es *executionState) do() {
	var wg sync.WaitGroup
	for _, src := range es.sources {
		wg.Add(1)
		go func(src Source) {
			ctx := es.ctx
			if ctxWithSpan, span := StartSpanFromContext(ctx, reflect.TypeOf(src).String(), src.Label()); span != nil {
				ctx = ctxWithSpan
				defer span.Finish()
			}
			defer wg.Done()

			// Setup panic handling on the source goroutines
			defer func() {
				if e := recover(); e != nil {
					// We had a panic, abort the entire execution.
					err, ok := e.(error)
					if !ok {
						err = fmt.Errorf("%v", e)
					}

					if errors.Code(err) == codes.ResourceExhausted {
						es.abort(err)
						return
					}

					err = errors.Wrap(err, codes.Internal, "panic")
					es.abort(err)
					if entry := es.logger.Check(zapcore.InfoLevel, "Execute source panic"); entry != nil {
						entry.Stack = string(debug.Stack())
						entry.Write(zap.Error(err))
					}
				}
			}()
			src.Run(ctx)

			if mdn, ok := src.(MetadataNode); ok {
				es.metaCh <- mdn.Metadata()
			}
		}(src)
	}

	wg.Add(1)
	//es.dispatcher.Start(es.resources.ConcurrencyQuota, es.ctx)
	go func() {
		defer wg.Done()

		// Wait for all transports to finish
		//for _, t := range es.transports {
		//	select {
		//	case <-t.Finished():
		//	case <-es.ctx.Done():
		//		es.abort(es.ctx.Err())
		//	case err := <-es.dispatcher.Err():
		//		if err != nil {
		//			es.abort(err)
		//		}
		//	}
		//}
		//// Check for any errors on the dispatcher
		//err := es.dispatcher.Stop()
		//if err != nil {
		//	es.abort(err)
		//}

		// Wait for all transports to finish
		for i := 0; i < len(es.transports); i++ {
			// wait for the end of one kind operator in all pipeline threads
			for j := 0; j < len(es.ESmultiThreadPipeLine); j++ {
				select {
				case <-es.ESmultiThreadPipeLine[j].Worker[i].Finished():
					//log.Println(es.ESmultiThreadPipeLine[j].Worker[i].Label(), " in pipe ", j, " Finished()")
				case <-es.ctx.Done():
					//log.Println("ex.ctx.Done()!")
					es.abort(es.ctx.Err())
				case err := <-es.ESmultiThreadPipeLine[j].Worker[i].worker.Err():
					if err != nil {
						es.abort(err)
					}
				}
			}
			// stop pipe worker
			StopAllOperatorThread(i)
		}

		//log.Println("all transports are done")

	}()

	go func() {
		defer close(es.metaCh)
		wg.Wait()
	}()
}

// Need a unique stream context per execution context
type executionContext struct {
	es            *executionState
	parents       []DatasetID
	streamContext streamContext
}

func resolveTime(qt flux.Time, now time.Time) Time {
	return Time(qt.Time(now).UnixNano())
}

func (ec executionContext) Context() context.Context {
	return ec.es.ctx
}

func (ec executionContext) ResolveTime(qt flux.Time) Time {
	return resolveTime(qt, ec.es.p.Now)
}

func (ec executionContext) StreamContext() StreamContext {
	return ec.streamContext
}

func (ec executionContext) Allocator() *memory.Allocator {
	return ec.es.alloc
}

func (ec executionContext) Parents() []DatasetID {
	return ec.parents
}
