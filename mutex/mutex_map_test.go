package mutex

import (
	"fmt"
	"sync"
	"testing"
)

func TestMutexMap(t *testing.T) {
	var wg sync.WaitGroup

	count := 10
	m := NewMutexMap()

	for i := range count {
		wg.Add(1)

		go func() {
			defer wg.Done()

			m.Set("key", fmt.Sprint(i))
			m.Delete("key")
		}()
	}

	wg.Wait()

	if m.Len() != 0 {
		t.Errorf("Mutex Map failed to safely write: %d", m.Len())
	}
}

func BenchmarkMutexMapWirteOperation(t *testing.B) {
	var wg sync.WaitGroup

	count := 100000
	m := NewMutexMap()

	for i := range count {
		wg.Add(1)

		go func() {
			defer wg.Done()

			m.Set("key", fmt.Sprint(i))
			m.Delete("key")
		}()
	}

	wg.Wait()
}

func BenchmarkSyncMapWirteOperation(t *testing.B) {
	var wg sync.WaitGroup
	var m sync.Map

	count := 100000

	for i := range count {
		wg.Add(1)

		go func() {
			defer wg.Done()

			m.Store("key", i)
			m.Delete("key")
		}()
	}

	wg.Wait()
}
