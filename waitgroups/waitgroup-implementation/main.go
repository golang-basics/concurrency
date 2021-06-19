package main

import (
	"fmt"
	"sync/atomic"
)

type waitGroup struct {
	counter int64
}

func (wg *waitGroup) Add(n int64) {
	atomic.AddInt64(&wg.counter, n)
}

func (wg *waitGroup) Done() {
	atomic.AddInt64(&wg.counter, -1)
	if atomic.LoadInt64(&wg.counter) < 0 {
		panic("negative wait group counter")
	}
}

func (wg *waitGroup) Wait() {
	for {
		if atomic.LoadInt64(&wg.counter) == 0 {
			return
		}
	}
}

func main() {
	var wg waitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		fmt.Println("go routine 1")
	}()
	go func() {
		defer wg.Done()
		fmt.Println("go routine 2")
	}()
	wg.Wait()
	fmt.Printf("all go routines are done")
}
