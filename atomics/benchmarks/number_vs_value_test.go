package benchmarks

import (
	"sync"
	"sync/atomic"
	"testing"
)

// to run the benchmarks make sure to cd into "benchmarks" and run
// go test -bench=.
func BenchmarkStoreInt64(b *testing.B) {
	var wg sync.WaitGroup
	var count int64
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			defer wg.Done()
			atomic.StoreInt64(&count, int64(i))
		}(i)
	}
	wg.Wait()
}

// to run the benchmarks make sure to cd into "benchmarks" and run
// go test -bench=.
func BenchmarkStoreValue(b *testing.B) {
	var wg sync.WaitGroup
	var v atomic.Value
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func(i int) {
			defer wg.Done()
			// In a real world it would take even more time
			// because you have to do type conversion
			// from interface{} to whatever you think it's stored there
			// something like the following
			// _ = v.Load().(int64)
			v.Store(int64(i))
		}(i)
	}
	wg.Wait()
}
