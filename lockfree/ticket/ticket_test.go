package ticket

import (
	"sync"
	"testing"
)

func TestLockFreeTicket(t *testing.T) {
	var sharedResource int
	var wg sync.WaitGroup

	count := 100000
	ticket := NewTicket()

	for range count {
		wg.Add(1)

		go func() {
			defer wg.Done()
			ticket.Lock()
			defer ticket.Unlock()
			sharedResource++
		}()
	}
	wg.Wait()

	if sharedResource != count {
		t.Errorf("Shared resource is not equal to 10000. Got %d", sharedResource)
	}
}

func BenchmarkLockFreeTicket(t *testing.B) {
	var sharedResource int
	var wg sync.WaitGroup

	count := 100000
	ticket := NewTicket()

	for range count {
		wg.Add(1)

		go func() {
			defer wg.Done()
			ticket.Lock()
			defer ticket.Unlock()
			sharedResource++
		}()
	}
	wg.Wait()
}

func BenchmarkMutexLock(t *testing.B) {
	var sharedResource int
	var wg sync.WaitGroup
	var mu sync.Mutex

	count := 100000

	for range count {
		wg.Add(1)

		go func() {
			defer wg.Done()
			mu.Lock()
			defer mu.Unlock()
			sharedResource++
		}()
	}
	wg.Wait()
}
