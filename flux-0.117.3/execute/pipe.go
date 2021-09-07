package execute

import (
	"context"
	"github.com/influxdata/flux"
	"log"
	"sync"
)

var OperatorMap map[string]*consecutiveTransport = make(map[string]*consecutiveTransport, 10)

//var	ByName map[string]int = make(map[string]int, 99)
//var	ByIndex map[int]*consecutiveTransport = make(map[int]*consecutiveTransport, 99)
//
//func FindNextConsecutiveTransport(operatorName string) *consecutiveTransport {
//	if ByName[operatorName] == len(ByName)-1 {
//		return nil
//	}
//	return ByIndex[ByName[operatorName]+1]
//}

type pipeWorker struct {

	message	chan flux.Table
	t Transformation
	closed  bool
	closing chan struct{}

	ctx context.Context
	mu sync.Mutex
}

func newPipeWorker(t Transformation) *pipeWorker {
	return &pipeWorker{
		// size of buffered channel is 10 temporarily
		message: make(chan flux.Table, 64),
		t: t,
		closed: false,
		closing: make(chan struct{}),
	}
}

func (p *pipeWorker) Start(ct *consecutiveTransport, ctx context.Context)  {
	go func() {
		defer func() {
			log.Println("pipe worker finished")
		}()
		p.run(ct, ctx)
	}()
}

func (p *pipeWorker) Stop() error {
	// Check if this is the first time invoking this method.
	p.mu.Lock()
	if !p.closed {
		// If not, mark the dispatcher as closed and signal to the current
		// workers that they should stop processing more work.
		p.closed = true
		close(p.closing)
	}
	p.mu.Unlock()

	log.Println("stop pipeWorker")

	return nil
}

func (p *pipeWorker) run(ct *consecutiveTransport,ctx context.Context)  {
	for  {
		select {
		case <-ctx.Done():
			return
		case <-p.closing:
			return
		case msg := <- p.message:
			pipeProcess(ctx, ct, msg)
			// send the result to next operator channel

		}
	}
}

