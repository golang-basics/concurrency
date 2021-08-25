package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var condition int32
	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		// the below for loop is inefficient, wastes too much time on sleeping
		// wasted extra 90 milliseconds
		defer wg.Done()
		for {
			time.Sleep(100 * time.Millisecond)
			if atomic.LoadInt32(&condition) == 1 {
				fmt.Println("go routine 1: done")
				return
			}
		}
	}()
	go func() {
		// the below for loop is inefficient, wastes too much time on sleeping
		// wasted extra 90 milliseconds
		defer wg.Done()
		for {
			time.Sleep(100 * time.Millisecond)
			if atomic.LoadInt32(&condition) == 1 {
				fmt.Println("go routine 2: done")
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		time.Sleep(10 * time.Millisecond)
		atomic.StoreInt32(&condition, 1)
	}()
	wg.Wait()
}
