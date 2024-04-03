package spinningcas

import (
	"runtime"
	"sync/atomic"
)

const free int32 = 0

// Spinning Compare And Swap Lock
type SpinningCas struct {
	state int32
}

func newCas() *SpinningCas {
	return &SpinningCas{}
}

func (t *SpinningCas) Lock() {
	for !atomic.CompareAndSwapInt32(&t.state, free, 1) {
		// Unlike mutex, other work can be done while waiting for lock to be free.
		runtime.Gosched()
	}
}

func (t *SpinningCas) Unlock() {
	atomic.StoreInt32(&t.state, free)
}
