package main

import (
	"fmt"
	"time"
)

// sometimes we never know when and if a channel is ready
// a good practice is to always have a fallback
// a fallback can be a timeout or a default case
// to prevent the select statement from blocking
// a default case works very well with an infinite loop
// allowing other work to be done, before a channel becomes ready
func main() {
	done := make(chan struct{})
	go func() {
		time.Sleep(5 * time.Second)
		close(done)
	}()

	work := 0
	defer func() {
		fmt.Println("achieved", work, "cycles of work")
	}()

	for {
		select {
		case <-done:
			fmt.Println("done working")
			return
		default:
			work++
			fmt.Println("working...")
			time.Sleep(time.Second)
		}
	}
}
