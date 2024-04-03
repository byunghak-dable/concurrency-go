package spinningcas

import (
	"sync"
	"testing"
)

func TestLockFreeTicket(t *testing.T) {
	var sharedResource int
	var wg sync.WaitGroup

	count := 100000
	cas := newCas()

	for range count {
		wg.Add(1)

		go func() {
			defer wg.Done()
			cas.Lock()
			defer cas.Unlock()
			sharedResource++
		}()
	}
	wg.Wait()

	if sharedResource != count {
		t.Errorf("Shared resource is not equal to 10000. Got %d", sharedResource)
	}
}

func BenchmarkSpinningCas(t *testing.B) {
	var sharedResource int
	var wg sync.WaitGroup

	count := 10000
	cas := newCas()

	for range count {
		wg.Add(1)

		go func() {
			defer wg.Done()
			cas.Lock()
			defer cas.Unlock()
			sharedResource++
		}()
	}
	wg.Wait()
}

func BenchmarkMutex(t *testing.B) {
	var sharedResource int
	var wg sync.WaitGroup
	var mu sync.Mutex

	count := 10000

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
