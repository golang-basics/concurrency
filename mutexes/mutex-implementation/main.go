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
	// this could quickly be detected as a race condition
	// due to too many go routines' status being awake
	// Also there's no way we can control the go routines queue
	// or have access to the runtime internals
	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			mu.Lock()
			count++
			mu.Unlock()
		}()
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
	for atomic.CompareAndSwapInt32(&mu.state, 1, 0) {
		return
	}
	panic("unlock of unlocked mutex")
}
