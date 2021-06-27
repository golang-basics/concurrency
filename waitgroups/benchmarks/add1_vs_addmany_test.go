package benchmarks

import (
	"sync"
	"testing"
)

// run all the benchmarks in the current directory
// cd benchmarks
// go test -bench=.
func BenchmarkAddOne(b *testing.B) {
	var count int
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
}

// run all the benchmarks in the current directory
// cd benchmarks
// go test -bench=.
func BenchmarkAddMany(b *testing.B) {
	var count int
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			count++
		}()
	}
	wg.Wait()
}
