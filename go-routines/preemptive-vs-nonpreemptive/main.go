package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done := make(chan struct{})
	time.AfterFunc(time.Second, func() {
		close(done)
	})

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		defer fmt.Println("preemptive go routine done")

		select {
		case <-done:
			return
		case <-time.After(5 * time.Second):
		}

		fmt.Println("work that is cancelled")
	}()
	go func() {
		defer wg.Done()
		defer fmt.Println("non-preemptive go routine done")

		select {
		case <-time.After(5 * time.Second):
		}

		fmt.Println("work that is not cancelled")
	}()

	wg.Wait()
}
