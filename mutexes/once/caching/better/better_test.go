package main

import (
	"sync"
	"testing"
	"time"
)

// cd "caching"
// go test -bench . ./...
func BenchmarkBetter(b *testing.B) {
	b.ReportAllocs()
	c := client{
		cache: map[string]*cacheEntry{},
		mu:    new(sync.RWMutex),
	}
	var wg sync.WaitGroup
	wg.Add(b.N)
	for i := 0; i < b.N; i++ {
		go func() {
			defer wg.Done()
			c.httpCall("req", 300*time.Millisecond)
		}()
	}
	wg.Wait()
}
