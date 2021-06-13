package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	timeout := 2 * time.Second
	heartbeat, results := work(done, timeout/2, 1, 2, 3)
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
		case <-heartbeat:
			fmt.Println("pulse")
		case <-time.After(timeout):
			fmt.Println("timed out")
			return
		}
	}
}

func work(done chan struct{}, interval time.Duration, numbers ...int) (chan struct{}, chan int) {
	heartbeat, out := make(chan struct{}), make(chan int)
	go func() {
		defer close(heartbeat)
		defer close(out)

		time.Sleep(2 * time.Second)

		pulse := time.Tick(interval)
	loop:
		for _, number := range numbers {
			time.Sleep(time.Second)
			for {
				select {
				case <-done:
					return
				case <-pulse:
					select {
					case heartbeat <- struct{}{}:
					default:
					}
				case out <- number:
					continue loop
				}
			}
		}
	}()
	return heartbeat, out
}
