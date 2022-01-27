package main

import (
	"testing"
	"time"
)

// go test -bench=.
// compare Benchmark_basicMutex_load with Benchmark_rwMutex_load
func Benchmark_basicMutex_load(b *testing.B) {
	b.ReportAllocs()
	mu := basicMutex{readSleepDuration: time.Nanosecond}
	mu.store(10)
	for i := 0; i < b.N; i++ {
		go mu.load()
		go mu.load()
		go mu.load()
	}
}

func Benchmark_basicMutex_store(b *testing.B) {
	b.ReportAllocs()
	mu := basicMutex{readSleepDuration: time.Nanosecond}
	for i := 0; i < b.N; i++ {
		go mu.store(i)
	}
}

func Benchmark_basicMutex_hybrid(b *testing.B) {
	b.ReportAllocs()
	mu := basicMutex{readSleepDuration: time.Nanosecond}
	for i := 0; i < b.N; i++ {
		go mu.load()
		go mu.load()
		go mu.load()
		go mu.store(i)
	}
}

// go test -bench=.
// compare BenchmarkRWMutex_Load with BenchmarkBasicMutex_Load
func Benchmark_rwMutex_load(b *testing.B) {
	b.ReportAllocs()
	mu := rwMutex{readSleepDuration: time.Nanosecond}
	mu.store(10)
	for i := 0; i < b.N; i++ {
		go mu.load()
		go mu.load()
		go mu.load()
	}
}

func Benchmark_rwMutex_store(b *testing.B) {
	b.ReportAllocs()
	mu := rwMutex{readSleepDuration: time.Nanosecond}
	for i := 0; i < b.N; i++ {
		go mu.store(i)
	}
}

func Benchmark_rwMutex_hybrid(b *testing.B) {
	b.ReportAllocs()
	mu := rwMutex{readSleepDuration: time.Nanosecond}
	for i := 0; i < b.N; i++ {
		go mu.load()
		go mu.load()
		go mu.load()
		go mu.store(i)
	}
}
