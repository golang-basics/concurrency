package builtinmap_vs_syncmap

import (
	"strconv"
	"testing"
)

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
