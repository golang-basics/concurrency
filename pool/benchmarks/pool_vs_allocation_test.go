package benchmarks

import (
	"sync"
	"testing"
)

// cd benchmarks
// go test -bench=.
func BenchmarkPool(b *testing.B) {
	var p sync.Pool
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			p.Put(1)
			p.Get()
		}
	})
}

func BenchmarkAllocation(b *testing.B) {
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			i := 0
			i = i
		}
	})
}
