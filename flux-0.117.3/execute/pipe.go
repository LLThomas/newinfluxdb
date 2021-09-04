package execute

import (
	"context"
	"log"
)

type pipeWorker struct {

	message	chan Message

	ct *consecutiveTransport

	ctx context.Context
}

func newPipeWorker(t *consecutiveTransport) *pipeWorker {
	return &pipeWorker{
		// size of buffered channel is 10 temporarily
		message: make(chan Message, 10),
		ct: t,
	}
}

func (p *pipeWorker) Start(ctx context.Context)  {
	go func() {
		defer func() {
			log.Println("pipe worker finished")
		}()
		p.run(ctx)
	}()
}

func (p *pipeWorker) run(ctx context.Context)  {
	for  {
		select {
		// add done channel

		case msg := <- p.message:
			pipeProcess(ctx, p.ct.t, msg)
			// send the result to next operator channel

		}
	}
}