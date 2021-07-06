package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type calculator struct {
	res atomic.Value
}

func newCalculator() calculator {
	c := calculator{}
	c.res.Store(float64(0))
	return c
}

func (c *calculator) add(n float64) {
	c.res.Store(c.result() + n)
}

func (c *calculator) sub(n float64) {
	c.res.Store(c.result() - n)
}

func (c *calculator) mul(n float64) {
	c.res.Store(c.result() * n)
}

func (c *calculator) div(n float64) {
	if n == 0 {
		panic("cannot divide by zero")
	}
	c.res.Store(c.result() / n)
}

func (c *calculator) result() float64 {
	r, ok := c.res.Load().(float64)
	if !ok {
		panic("operating with wrong type")
	}
	return r
}

// to test this program properly, make sure to run it with -race flag
// go run -race main.go
func main() {
	c := newCalculator()
	var wg sync.WaitGroup

	wg.Add(5)
	go func() {
		defer wg.Done()
		c.add(10)
	}()
	go func() {
		defer wg.Done()
		c.sub(5)
	}()
	go func() {
		defer wg.Done()
		c.div(3)
	}()
	go func() {
		defer wg.Done()
		c.mul(4)
	}()
	go func() {
		defer wg.Done()
		c.add(13)
	}()

	wg.Wait()
	fmt.Println("result", c.result())
}
