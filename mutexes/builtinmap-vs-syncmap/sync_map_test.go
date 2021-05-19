package builtinmap_vs_syncmap

import (
	"strconv"
	"sync"
	"testing"
)

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
