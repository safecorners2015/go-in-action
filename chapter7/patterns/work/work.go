// Example...
// Package work manages a pool of goroutines to perform work.
package work

import (
	"sync"
)

// Worker must be implemented by types that want to use
// the work pool.

type Worker interface {
	Task()
}

// Pool provides a pool of gorutines that can execute any Worker
// tasks taht are submitted.
type Pool struct {
	work chan Worker
	wg   sync.WaitGroup
}

// Newe creates a new work pool.
func New(maxGoroutines int) *Pool {
	p := Pool{
		work: make(chan Worker),
	}

	p.wg.Add(maxGoroutines)
	for i := 0; i < maxGoroutines; i++ {
		go func() {
			for w := range p.work {
				w.Task()
			}
			p.wg.Done()
		}()
	}

	return &p
}

// Run submits work to the pool
func (p *Pool) Run(w Worker) {
	p.work <- w
}

// Shutdown waits for all the gorutines to shudown.
func (p *Pool) Shutdown() {
	close(p.work)
	p.wg.Wait()
}
