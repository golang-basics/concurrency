package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	heartbeat, results := work(done, 1, 2, 3)
	// if we don't wait for the first heartbeat
	// we won't receive any values do to timeout
	<-heartbeat
	fmt.Println("now we can begin ranging")
	for {
		select {
		case res, ok := <-results:
			if !ok {
				return
			}
			fmt.Println("result", res)
		case <-time.After(time.Second):
			fmt.Println("timed out")
			return
		}
	}
}

func work(done chan struct{}, numbers ...int) (chan struct{}, chan int) {
	heartbeat, out := make(chan struct{}), make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(out)

		time.Sleep(2 * time.Second)

		for _, number := range numbers {
			select {
			case heartbeat <- struct{}{}:
			default:
			}

			select {
			case <-done:
				return
			case out <- number:
			}
		}
	}()
	return heartbeat, out
}
