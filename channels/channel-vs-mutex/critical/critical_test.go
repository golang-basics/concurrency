package critical

import (
	"strconv"
	"testing"
)

// to run the benchmarks, cd into "critical" directory and run:
// go test -bench=.
func BenchmarkCriticalMutexCache(b *testing.B) {
	cache := newMutexCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		go cache.set("key", "value"+strconv.Itoa(i))
	}
}

// to run the benchmarks, cd into "critical" directory and run:
// go test -bench=.
func BenchmarkCriticalChannelCache(b *testing.B) {
	cache := newChannelCache()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		go cache.set("key", "value"+strconv.Itoa(i))
	}
}
