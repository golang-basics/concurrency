package benchmarks

import (
	"context"
	"fmt"
	"sync"
	"testing"
	"time"
)

const timeout = 100 * time.Millisecond

// To run the benchmarks make sure to cd into "patterns/context/context-vs-done-channel/benchmarks"
// go test -bench=.
func BenchmarkDoneChannel(b *testing.B) {
	done := make(chan struct{})

	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go requestWithChannel(done, &wg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(timeout)
		close(done)
	}()

	wg.Wait()
}

// To run the benchmarks make sure to cd into "patterns/context/context-vs-done-channel/benchmarks"
// go test -bench=.
func BenchmarkContext(b *testing.B) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)

	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go requestWithContext(ctx, &wg)
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		time.Sleep(timeout)
		cancel()
	}()

	wg.Wait()
}

func requestWithChannel(done <-chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	select {
	case <-done:
		return
	case <-time.After(time.Second):
		fmt.Println("this always times out")
	}
}

func requestWithContext(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	if deadline, ok := ctx.Deadline(); ok {
		if deadline.Sub(time.Now().Add(timeout)) <= 0 {
			return
		}
	}
	select {
	case <-ctx.Done():
		return
	case <-time.After(time.Second):
		fmt.Println("this always times out")
	}
}
