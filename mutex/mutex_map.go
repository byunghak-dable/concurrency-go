package mutex

import "sync"

type MutexMap struct {
	valueByKey map[string]string
	mu         sync.RWMutex
}

func NewMutexMap() MutexMap {
	return MutexMap{
		valueByKey: make(map[string]string),
	}
}

func (m *MutexMap) Set(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	m.valueByKey[key] = value
}

func (m *MutexMap) Delete(key string) {
	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.valueByKey, key)
}

func (m *MutexMap) Get(key string) string {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return m.valueByKey[key]
}

func (m *MutexMap) Len() int {
	m.mu.RLock()
	defer m.mu.RUnlock()

	return len(m.valueByKey)
}
