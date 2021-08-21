package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// sync.Once implementation
// https://github.com/golang/go/blob/master/src/sync/once.go#L14
type once struct {
	done uint32
	mu   sync.Mutex
}

// sync.Once.Do implementation
// https://github.com/golang/go/blob/master/src/sync/once.go#L42
func (o *once) Do(fn func()) {
	if atomic.LoadUint32(&o.done) == 0 {
		o.mu.Lock()
		defer o.mu.Unlock()
		if o.done == 0 {
			defer atomic.StoreUint32(&o.done, 1)
			fn()
		}
	}
}

// try running this with the -race flag
// go run -race main.go
func main() {
	var o once
	f := func(i int) func() {
		return func() {
			fmt.Println("printing once:", i)
		}
	}

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		o.Do(f(1))
	}()
	go func() {
		defer wg.Done()
		o.Do(f(2))
	}()
	go func() {
		defer wg.Done()
		o.Do(f(3))
	}()

	wg.Wait()
}
