package benchmarks

import (
	"testing"
)

// compare BenchmarkBasicMutex_Load with BenchmarkRWMutex_Load
func BenchmarkBasicMutex_Load(b *testing.B) {
	mu := BasicMutex{}
	mu.Store(10)
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
	}
}

func BenchmarkBasicMutex_Store(b *testing.B) {
	mu := BasicMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Store(i)
	}
}

func BenchmarkBasicMutex_Hybrid(b *testing.B) {
	mu := BasicMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
		go mu.Store(i)
	}
}

// compare BenchmarkRWMutex_Load with BenchmarkBasicMutex_Load
func BenchmarkRWMutex_Load(b *testing.B) {
	mu := RWMutex{}
	mu.Store(10)
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
	}
}

func BenchmarkRWMutex_Store(b *testing.B) {
	mu := RWMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Store(i)
	}
}

func BenchmarkRWMutex_Hybrid(b *testing.B) {
	mu := RWMutex{}
	for i := 0; i < b.N; i++ {
		go mu.Load()
		go mu.Load()
		go mu.Load()
		go mu.Store(i)
	}
}
