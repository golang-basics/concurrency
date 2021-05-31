package main

import (
	"testing"
)

// to run the benchmarks, cd into the stream-vs-batch directory and run
// go test -bench=. -benchtime=3s
func BenchmarkStreamPipeline(b *testing.B) {
	b.ReportAllocs()
	numbers := gen(1_000_000)
	for i := 0; i < b.N; i++ {
		for _, num := range numbers {
			sMultiply(sAdd(sMultiply(num, 2), 1), 2)
		}
	}
}

// to run the benchmarks, cd into the stream-vs-batch directory and run
// go test -bench=. -benchtime=3s
func BenchmarkBatchPipeline(b *testing.B) {
	numbers := gen(1_000_000)
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		for range bMultiply(bAdd(bMultiply(numbers, 2), 1), 2) {
		}
	}
}
