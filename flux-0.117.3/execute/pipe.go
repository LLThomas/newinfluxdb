package execute

import (
	"context"
	"github.com/influxdata/flux"
	"log"
	"sync"
)

var OperatorMap map[string]*consecutiveTransport = make(map[string]*consecutiveTransport, 10)
var ResOperator Transformation

func ConnectOperator(name string, b flux.Table) {

	log.Println("ConnectOperator: ")

	next := OperatorMap[name]
	if OperatorMap[name] == nil {
		// res handler
		log.Println("res: ", b.Key())
		ResOperator.Process(DatasetID{0}, b)
	} else {
		log.Println("next: ", next.Label())
		next.PushToChannel(b)
	}
	log.Println(999)
}

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
			// no error handler , no //case err := <-es.dispatcher.Err(): in executor.go: TODO
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
			ct.pipeProcesses(ctx, msg)
		}
	}
}

