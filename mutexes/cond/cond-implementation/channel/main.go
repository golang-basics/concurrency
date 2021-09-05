package main

import "sync"

func main() {
}

type cond struct {
	L  sync.Mutex // used by caller
	mu sync.Mutex // guards ch
	ch chan struct{}
}

func (c *cond) Wait() {
	c.mu.Lock()
	ch := c.ch
	c.mu.Unlock()
	c.L.Unlock()
	<-ch
	c.L.Lock()
}

func (c *cond) Signal() {
	c.mu.Lock()
	defer c.mu.Unlock()
	select {
	case c.ch <- struct{}{}:
	default:
	}
}

func (c *cond) Broadcast() {
	c.mu.Lock()
	defer c.mu.Unlock()
	close(c.ch)
	c.ch = make(chan struct{})
}
