package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	m := &sync.Mutex{}
	readersCond := sync.NewCond(m)
	validatorsCond := sync.NewCond(m)

	wg.Add(7)
	// writer
	go func() {
		fmt.Println("writer")
		for i :=0; i<3;i++ {
			validatorsCond.Signal()
			time.Sleep(100*time.Millisecond)
		}
		// the below call to Broadcast has enough time
		// so that the go routines rendezvous is ready
		// so that we avoid deadlocks
		readersCond.Broadcast()
		wg.Done()
	}()

	// validators
	go func() {
		validator(validatorsCond, 1)
		wg.Done()
	}()
	go func() {
		validator(validatorsCond, 2)
		wg.Done()
	}()
	go func() {
		validator(validatorsCond, 3)
		wg.Done()
	}()

	// readers
	go func() {
		reader(readersCond, 1)
		wg.Done()
	}()
	go func() {
		reader(readersCond, 2)
		wg.Done()
	}()
	go func() {
		reader(readersCond, 3)
		wg.Done()
	}()

	wg.Wait()
}

func reader(c *sync.Cond, i int) {
	c.L.Lock()
	defer c.L.Unlock()
	c.Wait()
	fmt.Println("reader", i)
}

func validator(c *sync.Cond, i int) {
	c.L.Lock()
	defer c.L.Unlock()
	c.Wait()
	fmt.Println("validator", i)
}
