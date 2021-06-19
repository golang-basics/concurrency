package benchmarks

import (
	"sync"
	"time"
)

type BasicMutex struct {
	mu    sync.Mutex
	value int
}

func (m *BasicMutex) Store(value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}

func (m *BasicMutex) Load() int {
	m.mu.Lock()
	time.Sleep(100 * time.Nanosecond)
	defer m.mu.Unlock()
	return m.value
}

type RWMutex struct {
	mu    sync.RWMutex
	value int
}

func (m *RWMutex) Store(value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}

func (m *RWMutex) Load() int {
	m.mu.RLock()
	time.Sleep(100 * time.Nanosecond)
	defer m.mu.RUnlock()
	return m.value
}
