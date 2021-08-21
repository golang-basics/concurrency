package mutex_vs_atomic

import (
	"sync"
	"sync/atomic"
	"testing"
)

// to run the benchmarks cd into "atomics/benchmarks" and run:
// go test -bench=.
func BenchmarkAtomicNumber(b *testing.B) {
	b.ReportAllocs()
	var wg sync.WaitGroup
	var v int64
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				atomic.SwapInt64(&v, int64(j))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	// to avoid compile errors
	v = v
}

// to run the benchmarks cd into "atomics/benchmarks" and run:
// go test -bench=.
func BenchmarkAtomicStruct(b *testing.B) {
	b.ReportAllocs()
	var wg sync.WaitGroup
	var v atomic.Value
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				v.Store(Config{
					a: []int{j + 1, j + 2, j + 3, j + 4, j + 5},
				})
			}
			wg.Done()
		}()
	}
	wg.Wait()
	// to avoid compile errors
	v = v
}
