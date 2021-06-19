package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	c := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		timeout := time.NewTimer(2 * time.Second)
		for {
			select {
			case <-timeout.C:
				fmt.Println("timeout")
				wg.Done()
				return
			case v, ok := <-c:
				if ok {
					fmt.Println("value:", v)
				} else {
					wg.Done()
					return
				}
			}
		}
	}()
	c <- 1
	time.Sleep(500 * time.Millisecond)
	close(c)
	wg.Wait()
}
