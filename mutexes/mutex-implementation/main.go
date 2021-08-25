package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// try running this example with the -race flag
// cd into "mutex-implementation"
// go run -race main.go
func main() {
	var count int
	var mu mutex
	var wg sync.WaitGroup

	// if we increase the number of go routines
	// this will quickly start to fail
	// due to go routines status being awake
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			mu.Lock()
			count = i
			mu.Unlock()
		}(i)
	}
	wg.Wait()

	fmt.Println("count", count)
}

type mutex struct {
	state int32
}

func (mu *mutex) Lock() {
	if atomic.CompareAndSwapInt32(&mu.state, 0, 1) {
		return
	}
	for {
		// continue
		atomic.AddInt32(&mu.state, 1)
		s := atomic.LoadInt32(&mu.state)
		if s > 1 {
			panic("all goroutines are asleep - deadlock!")
		}
		if s == 1 {
			return
		}
	}
}

func (mu *mutex) Unlock() {
	if atomic.CompareAndSwapInt32(&mu.state, 1, 0) {
		return
	}
	panic("unlock of unlocked mutex")
}
