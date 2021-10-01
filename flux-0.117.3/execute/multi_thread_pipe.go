package execute

import (
	"context"
	"github.com/influxdata/influxdb/v2/tsdb/cursors"
	"log"
	"sync"
)

// global wg (for done in reader.go:239)
var WG sync.WaitGroup

// global executionState
var ExecutionState *executionState

// global transformation operators line
var OperatorIndex map[string]int = make(map[string]int)
var OperatorMap map[string]*consecutiveTransport = make(map[string]*consecutiveTransport)
var ResOperator Transformation

type MultiThreadPipeLine struct {
	// series key
	DataSource	[]cursors.Cursor
	Current		[]cursors.Cursor

	// worker in this pipeline
	Worker 		[]*consecutiveTransport
}

func newMultiPipeLine(dataSource []cursors.Cursor, current []cursors.Cursor, worker []*consecutiveTransport) *MultiThreadPipeLine {
	mpl := &MultiThreadPipeLine{
		DataSource: dataSource,
		Current: current,
		Worker: worker,
	}
	return mpl
}

func (mpl *MultiThreadPipeLine) startPipeLine(ctx context.Context)  {
	//log.Println(mpl.Worker)
	for i := 0; i < len(mpl.Worker); i++ {
		//log.Println("start pipe worker: ", mpl.Worker[i].Label())
		mpl.Worker[i].startPipeWorker(ctx)
	}
}

//func (mpl *MultiThreadPipeLine) stopPipeLine() error {
//	for i := 0; i < len(mpl.Worker); i++ {
//		log.Println("stop pipe worker: ", mpl.Worker[i].Label())
//		if err := mpl.Worker[i].worker.Stop(); err != nil {
//			return err
//		}
//	}
//	return nil
//}

func StopAllOperatorThread(whichOperator int) error {
	mpl := ExecutionState.ESmultiThreadPipeLine
	for i := 0; i < len(mpl); i++ {
		log.Println("stop operator: ", mpl[i].Worker[whichOperator].Label(), " in ", i)
		if  err := mpl[i].Worker[whichOperator].worker.Stop(); err != nil {
			return err
		}
	}
	return nil
}

// split series key for each pipeline
func SplitSeriesKey(allSeriesKey []cursors.Cursor) map[cursors.Cursor]int {

	pipeToGroupKey := make(map[cursors.Cursor]int)

	mpl := ExecutionState.ESmultiThreadPipeLine
	n := len(allSeriesKey)
	split := n / len(mpl)
	i := 0
	for len(allSeriesKey) > 0 {
		count := 0
		for count < split && len(allSeriesKey) > 0 {

			pipeToGroupKey[allSeriesKey[0]] = i

			mpl[i%len(mpl)].DataSource = append(mpl[i%len(mpl)].DataSource, allSeriesKey[0])
			allSeriesKey = allSeriesKey[1:]
			count++
		}
		i++
	}
	return pipeToGroupKey
}

// dispatch datasource to current and send to first operator
func DispatchAndSend(ff func(whichPipeLine int))  {
	mpl := ExecutionState.ESmultiThreadPipeLine
	for i := 0; i < len(mpl); i++ {
		for len(mpl[i].DataSource) > 0 {
			// set the size of current to 3 (executor.go:139 make([]cursors.Cursor, 3))
			// fill the current with DataSource (first time)
			//for j := 0; j < len(mpl[i].Current) && len(mpl[i].DataSource) > 0; j++ {
			//	mpl[i].Current[j] = mpl[i].DataSource[0]
			//	mpl[i].DataSource = mpl[i].DataSource[1:]
			//}
			ff(i)
		}
	}
}