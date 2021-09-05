package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// c := sync.NewCond(new(sync.Mutex))
	c := newCond(new(sync.Mutex))
	var wg sync.WaitGroup

	wg.Add(3)
	go func() {
		defer wg.Done()
		c.L.Lock()
		c.Wait()
		fmt.Println("go routine 1")
		c.L.Unlock()
	}()
	go func() {
		defer wg.Done()
		c.L.Lock()
		c.Wait()
		fmt.Println("go routine 2")
		c.L.Unlock()
	}()
	go func() {
		defer wg.Done()
		c.L.Lock()
		c.Wait()
		fmt.Println("go routine 3")
		c.L.Unlock()
	}()

	time.Sleep(time.Second)
	// c.Signal()
	// c.Signal()
	// c.Signal()
	c.Broadcast()
	wg.Wait()
}

func newCond(lock sync.Locker) cond {
	return cond{L: lock}
}

// Taken from the source code from standard library sync.Cond
// ----------------------------------------------------------
// Cond implements a condition variable, a rendezvous point
// for goroutines waiting for or announcing the occurrence
// of an event.
//
// Each Cond has an associated Locker L (often a *Mutex or *RWMutex),
// which must be held when changing the condition and
// when calling the Wait method.
//
// A Cond must not be copied after first use.
type cond struct {
	wait   int32
	notify int32
	L      sync.Locker
}

func (c *cond) Wait() {
	t := atomic.AddInt32(&c.wait, 1) - 1
	c.L.Unlock()
	c.notifyListWait(t)
	c.L.Lock()
}

func (c *cond) Signal() {
	if atomic.LoadInt32(&c.notify) == atomic.LoadInt32(&c.wait) {
		return
	}
	atomic.AddInt32(&c.notify, 1)
}

func (c *cond) Broadcast() {
	if atomic.LoadInt32(&c.notify) == atomic.LoadInt32(&c.wait) {
		return
	}
	atomic.StoreInt32(&c.notify, atomic.LoadInt32(&c.wait))
}

func (c *cond) notifyListWait(t int32) {
	for {
		if atomic.LoadInt32(&c.notify) > t {
			return
		}
	}
}
