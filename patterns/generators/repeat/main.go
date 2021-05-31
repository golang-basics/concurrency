package main

import (
	"concurrency/patterns/generators"
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go func() {
		select {
		case <-time.After(3 * time.Second):
			close(done)
		}
	}()

	for num := range generators.Repeat(done, 1, 2, 3) {
		fmt.Printf("%d ", num)
		time.Sleep(50 * time.Millisecond)
	}
}
