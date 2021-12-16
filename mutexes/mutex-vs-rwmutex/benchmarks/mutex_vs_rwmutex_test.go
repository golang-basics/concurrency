package benchmarks

import (
	"testing"
)

// go test -bench=.
// compare BenchmarkBasicMutex_Load with BenchmarkRWMutex_Load
func BenchmarkBasicMutex_Load(b *testing.B) {
	b.ReportAllocs()
	mu := BasicMutex{}
	mu.Store(10)
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
	}
}

func BenchmarkBasicMutex_Store(b *testing.B) {
	b.ReportAllocs()
	mu := BasicMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Store(i)
	}
}

func BenchmarkBasicMutex_Hybrid(b *testing.B) {
	b.ReportAllocs()
	mu := BasicMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
		go mu.Store(i)
	}
}

// go test -bench=.
// compare BenchmarkRWMutex_Load with BenchmarkBasicMutex_Load
func BenchmarkRWMutex_Load(b *testing.B) {
	b.ReportAllocs()
	mu := RWMutex{}
	mu.Store(10)
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
	}
}

func BenchmarkRWMutex_Store(b *testing.B) {
	b.ReportAllocs()
	mu := RWMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Store(i)
	}
}

func BenchmarkRWMutex_Hybrid(b *testing.B) {
	b.ReportAllocs()
	mu := RWMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
		go mu.Store(i)
	}
}
