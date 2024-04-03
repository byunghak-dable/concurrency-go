package spinningcas

import "sync/atomic"

// Spinning Compare And Swap
type SpinningCas struct {
	current int32
	next    int32
}

func newCas() *SpinningCas {
	return &SpinningCas{}
}

func (t *SpinningCas) Lock() {
	myTicket := atomic.AddInt32(&t.next, 1) - 1

	for atomic.LoadInt32(&t.current) != myTicket {
	}
}

func (t *SpinningCas) Unlock() {
	atomic.AddInt32(&t.current, 1)
}
