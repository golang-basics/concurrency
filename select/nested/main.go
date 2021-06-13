package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	done, ping := make(chan struct{}), make(chan struct{})
	results := make(chan int)
	defer close(done)

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		select {
		case <-done:
			return
		case <-ping:
			fmt.Println("pong")
		}
	}()

	go func() {
		defer wg.Done()
		select {
		case <-done:
			return
		case results <- 1:
			select {
			case ping <- struct{}{}:
				fmt.Println("ping")
			default:
				fmt.Println("default")
			}
		}
	}()

	select {
	case <-done:
		return
	case r := <-results:
		fmt.Println("result", r)
	case <-time.After(time.Second):
		fmt.Println("timed out")
	}

	wg.Wait()
}
