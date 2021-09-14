package execute

import (
	"context"
	"fmt"
	"github.com/influxdata/flux"
	"github.com/influxdata/flux/codes"
	"github.com/influxdata/flux/internal/errors"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"log"
	"runtime/debug"
	"sync"
)

var OperatorMap map[string]*consecutiveTransport = make(map[string]*consecutiveTransport, 10)
var ResOperator Transformation

type pipeWorker struct {

	message	chan flux.Table
	t 		Transformation
	closed  bool
	closing chan struct{}
	wg 		sync.WaitGroup

	ctx 	context.Context
	mu 		sync.Mutex
	err 	error
	errC	chan error

	logger 	*zap.Logger
}

func newPipeWorker(t Transformation, logger *zap.Logger) *pipeWorker {
	return &pipeWorker{
		message: make(chan flux.Table, 100),
		t: t,
		closed: false,
		closing: make(chan struct{}),
		errC: make(chan error, 1),
		logger: logger.With(zap.String("component", "pipeWorker")),
	}
}

func (p *pipeWorker) Start(ct *consecutiveTransport, ctx context.Context)  {
	p.wg.Add(1)
	go func() {
		defer p.wg.Done()
		defer func() {
			// no error handler , no //case err := <-es.dispatcher.Err(): in executor.go: TODO
			//log.Println("pipe worker finished")
			if e := recover(); e != nil {
				log.Println("pipe worker error: ", e)
				err, ok := e.(error)
				if !ok {
					err = fmt.Errorf("%v", e)
				}

				if errors.Code(err) == codes.ResourceExhausted {
					p.setErr(err)
					return
				}

				err = errors.Wrap(err, codes.Internal, "panic")
				p.setErr(err)
				if entry := p.logger.Check(zapcore.InfoLevel, "Dispatcher panic"); entry != nil {
					entry.Stack = string(debug.Stack())
					entry.Write(zap.Error(err))
				}
			}
		}()
		p.run(ct, ctx)
	}()
}

func (p *pipeWorker) Err() <-chan error {
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.errC
}

func (p *pipeWorker) setErr(err error) {
	p.mu.Lock()
	defer p.mu.Unlock()
	if p.err == nil {
		p.err = err
		p.errC <- err
	}
}

func (p *pipeWorker) Stop() error {

	//log.Println("stop pipeWorker")

	// Check if this is the first time invoking this method.
	p.mu.Lock()
	if !p.closed {
		// If not, mark the dispatcher as closed and signal to the current
		// workers that they should stop processing more work.
		p.closed = true
		close(p.closing)
	}
	p.mu.Unlock()

	// Wait for the existing workers to finish.
	p.wg.Wait()

	// Grab the error from within a lock.
	p.mu.Lock()
	defer p.mu.Unlock()
	return p.err
}

func (p *pipeWorker) run(ct *consecutiveTransport,ctx context.Context)  {
	for  {
		select {
		case <-ctx.Done():
			//log.Println("ctx.Done()")
			return
		case <-p.closing:
			//log.Println("p.closing")
			return
		case msg := <- p.message:
			//log.Println("pipeWorker msg")
			ct.pipeProcesses(ctx, msg)
		}
	}
}

