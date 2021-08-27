package benchmarks

import (
	"sync"
	"testing"
)

// cd into benchmarks
// go test -bench .
func BenchmarkSignalCond(b *testing.B) {
	b.ReportAllocs()
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			cond.L.Lock()
			defer cond.L.Unlock()
			wg.Done()
			cond.Wait()
		}()
	}
	for i := 0; i < b.N; i++ {
		cond.Signal()
	}
	wg.Wait()
}

func BenchmarkBroadcastCond(b *testing.B) {
	b.ReportAllocs()
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			cond.L.Lock()
			defer cond.L.Unlock()
			wg.Done()
			cond.Wait()
		}()
	}
	cond.Broadcast()
	wg.Wait()
}

func BenchmarkSignalChannel(b *testing.B) {
	b.ReportAllocs()
	signal := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-signal
		}()
	}
	for i := 0; i < b.N; i++ {
		signal <- struct{}{}
	}
	wg.Wait()
}

func BenchmarkBroadcastChannel(b *testing.B) {
	b.ReportAllocs()
	broadcast := make(chan struct{})
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			<-broadcast
		}()
	}

	close(broadcast)
	wg.Wait()
}
