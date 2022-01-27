// cd benchmarks
// go test -bench=. -benchtime=3s ./...
package benchmarks

import (
	"sync"
	"testing"
)

// Because this is a relatively simple example using a very small memory footprint
// with only 1 operation, it will outperform the other 2 benchmarks most of the time.
// However, if we start playing with more memory and operations, things will start paying off for the mutex example
func BenchmarkSimpleCounter(b *testing.B) {
	b.ReportAllocs()
	var count int
	for i := 0; i < b.N; i++ {
		count++
	}
}

// This one is the most inefficient, due to the unnecessary wait time between each iteration
// Essentially making it perform even worse than a regular sequential execution
func BenchmarkWaitGroupCounter(b *testing.B) {
	b.ReportAllocs()
	var count int
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			count++
			wg.Done()
		}()
		wg.Wait()
	}
}

// We're using a bench time to avoid using a WaitGroup for better benchmark precision
// go test -bench=. -benchtime=3s ./...
func BenchmarkMutexCounter(b *testing.B) {
	b.ReportAllocs()
	var count int
	var mu sync.Mutex
	for i := 0; i < b.N; i++ {
		go func() {
			mu.Lock()
			count++
			mu.Unlock()
		}()
	}
}
