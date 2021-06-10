package benchmarks

import (
	"fmt"
	"sync"
	"testing"
)

// utilize and explain the Little Law
// L = λ*W
// L - average number of units in the system
// λ - average arrival rate of units
// W - average time a unit spends in the system

var (
	bufferedWrites   int
	unbufferedWrites int
)

func TestMain(m *testing.M) {
	m.Run()
	fmt.Println("buffered writes:", bufferedWrites)
	fmt.Println("un-buffered writes:", unbufferedWrites)
}

// to run the benchmark cd into "buffered-vs-unbuffered" and run:
// go test -bench=.
func BenchmarkUnbufferedChannel(b *testing.B) {
	b.ReportAllocs()
	done := make(chan struct{})
	defer close(done)
	ch := make(chan int)
	go reader(done, ch)

	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			unbufferedWrites++
			ch <- i
		}(i)
	}
	wg.Wait()
}

// to run the benchmark cd into "buffered-vs-unbuffered" and run:
// go test -bench=.
func BenchmarkBufferedChannel(b *testing.B) {
	b.ReportAllocs()
	done := make(chan struct{})
	defer close(done)

	bufferSize := 5000
	ch := make(chan int, bufferSize)
	go reader(done, ch)

	var size int
	if b.N-bufferSize < 0 {
		size = b.N
	} else {
		size = bufferSize
	}

	var wg sync.WaitGroup
	for i := 0; i < b.N; i += size {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			for j := 0; j < size; j++ {
				bufferedWrites++
				ch <- i + j
			}
		}(i)
	}
	wg.Wait()
}

func reader(done chan struct{}, in <-chan int) {
	for {
		select {
		case <-done:
			return
		case <-in:
		}
	}
}
