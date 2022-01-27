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
	if atomic.LoadUint32(&o.done) == 1 {
		return
	}

	// Slow-path.
	o.mu.Lock()
	defer o.mu.Unlock()
	if o.done == 0 {
		defer atomic.StoreUint32(&o.done, 1)
		fn()
	}
}

// https://github.com/matryer/resync
func (o *once) Reset() {
	o.mu.Lock()
	defer o.mu.Unlock()
	atomic.StoreUint32(&o.done, 0)
}

// try running this with the -race flag
// go run -race main.go
func main() {
	var i int
	var o once
	add := func(n int) func() {
		return func() {
			i += n
		}
	}

	o.Do(add(10))
	o.Reset()
	o.Do(add(-5))
	o.Do(add(100))

	fmt.Println("i:", i)
}
