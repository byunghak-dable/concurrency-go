package ticket

import "sync/atomic"

type Ticket struct {
	current int32
	next    int32
}

func NewTicket() *Ticket {
	return &Ticket{}
}

func (t *Ticket) Lock() {
	myTicket := atomic.AddInt32(&t.next, 1) - 1

	for atomic.LoadInt32(&t.current) != myTicket {
	}
}

func (t *Ticket) Unlock() {
	atomic.AddInt32(&t.current, 1)
}
