package benchmarks

import (
	"strconv"
	"sync"
	"testing"
)

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

func BenchmarkBuiltinStringMap_Load(b *testing.B) {
	m := NewBuiltinStringMap()
	m.Store("k1", "v1")
	for i := 0; i < b.N; i++ {
		go m.Load("k1")
	}
}

func BenchmarkBuiltinStringMap_Store(b *testing.B) {
	m := NewBuiltinStringMap()
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(i)
		go m.Store("k"+val, "v"+val)
	}
}

func BenchmarkBuiltinStringMap_Hybrid(b *testing.B) {
	m := NewBuiltinStringMap()
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(i)
		go m.Store("k"+val, "v"+val)
		go m.Load("k" + val)
	}
}

// 1 write and read heavy => better than built in map
func BenchmarkSyncMap_Load(b *testing.B) {
	m := sync.Map{}
	m.Store("k1", "v1")
	for i := 0; i < b.N; i++ {
		go m.Load("k1")
	}
}

// write/read only use a regular/built in map
func BenchmarkSyncMap_Store(b *testing.B) {
	m := sync.Map{}
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(i)
		go m.Store("k"+val, "v"+val)
	}
}

// write/read heavy with a lot of overwrites => better than built in map
func BenchmarkSyncMap_Hybrid(b *testing.B) {
	m := NewBuiltinStringMap()
	for i := 0; i < b.N; i++ {
		val := strconv.Itoa(i)
		go m.Store("k"+val, "v"+val)
		go m.Load("k" + val)
	}
}
