package ticket

import (
	"runtime"
	"sync/atomic"
)

type TicketStorage struct {
	slots  []string
	ticket uint64
	done   uint64
}

func NewTicketStorage() *TicketStorage {
	return &TicketStorage{}
}

func (t *TicketStorage) Put(s string) {
	ticket := atomic.AddUint64(&t.ticket, 1) - 1

	t.slots[ticket] = s

	for !atomic.CompareAndSwapUint64(&t.done, ticket, ticket+1) {
		runtime.Gosched()
	}
}

func (t *TicketStorage) GetDone() []string {
	return t.slots[:atomic.LoadUint64(&t.done)+1]
}
