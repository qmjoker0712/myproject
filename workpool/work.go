package workpool

import (
	"sync"
)

type Worker interface {
	Task() error
}

type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

func New(maxGoroutineNum int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutineNum)
	for i := 0; i < maxGoroutineNum; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

func (p *Pool) SubmitWork(w Worker) {
	p.work <- w
}

func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
