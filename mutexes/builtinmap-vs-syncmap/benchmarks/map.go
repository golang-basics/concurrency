package benchmarks

import "sync"

type BuiltinStringMap struct {
	internal map[string]string
	mu       sync.RWMutex
}

func NewBuiltinStringMap() BuiltinStringMap {
	return BuiltinStringMap{internal: map[string]string{}}
}

func (m *BuiltinStringMap) Load(key string) (string, bool) {
	m.mu.RLock()
	defer m.mu.RUnlock()
	v, ok := m.internal[key]
	return v, ok
}

func (m *BuiltinStringMap) Store(key, value string) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.internal[key] = value
}
