package main

import (
	"sync"
	"testing"
	"time"
)

// to run the benchmarks, cd into "channel-vs-waitgroup" directory and run:
// go test -bench=.
func BenchmarkWaitGroup(b *testing.B) {
	b.ReportAllocs()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			time.Sleep(2 * time.Nanosecond)
			wg.Done()
		}()
	}
	wg.Wait()
}

// to run the benchmarks, cd into "channel-vs-waitgroup" directory and run:
// go test -bench=.
func BenchmarkChannel(b *testing.B) {
	b.ReportAllocs()
	done := make(chan struct{})
	for i := 0; i < b.N; i++ {
		go func() {
			time.Sleep(2 * time.Nanosecond)
			done <- struct{}{}
		}()
	}
	for i := 0; i < b.N; i++ {
		<-done
	}
}
