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
		defer wg.Done()
		// the below for loop makes the CPU consume all cycles on 1 CORE
		for {
			if atomic.LoadInt32(&condition) == 1 {
				fmt.Println("go routine 1: done")
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		// the below for loop makes the CPU consume all cycles on 1 CORE
		for {
			if atomic.LoadInt32(&condition) == 1 {
				fmt.Println("go routine 2: done")
				return
			}
		}
	}()
	go func() {
		defer wg.Done()
		time.Sleep(time.Second)
		atomic.StoreInt32(&condition, 1)
	}()
	wg.Wait()
}
