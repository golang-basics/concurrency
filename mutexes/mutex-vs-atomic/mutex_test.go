package mutex_vs_atomic

import (
	"sync"
	"testing"
)

// to run the benchmarks cd into "atomics/benchmarks" and run:
// go test -bench=.
func BenchmarkMutexNumber(b *testing.B) {
	b.ReportAllocs()
	var wg sync.WaitGroup
	var mu sync.Mutex
	var v int64
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				mu.Lock()
				v = int64(j)
				mu.Unlock()
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
func BenchmarkMutexStruct(b *testing.B) {
	b.ReportAllocs()
	var wg sync.WaitGroup
	var mu sync.Mutex
	cfg := Config{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				mu.Lock()
				cfg = Config{
					a: []int{j + 1, j + 2, j + 3, j + 4, j + 5},
				}
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	// to avoid compile errors
	cfg = cfg
}
