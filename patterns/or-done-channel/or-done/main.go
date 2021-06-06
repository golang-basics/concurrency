// The Or-Done channel pattern consumes a channel till it's closed or till done is closed
// It's a very handy pattern because it ensures no further values from the channel will be consumed.
// By having 2 select statements and an early return on done channel
// we make sure to interrupt the channel consumption as soon as the done channel is closed

package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	// try to close the channel right away
	// you'll see no values being produced
	// close(done)
	go func() {
		time.Sleep(100 * time.Millisecond)
		close(done)
	}()
	for val := range orDone(done, gen(1, 2, 3)) {
		fmt.Println("val:", val)
	}
}

func orDone(done chan struct{}, in <-chan int) <-chan int {
	valStream := make(chan int)
	go func() {
		defer close(valStream)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if !ok {
					return
				}
				select {
				case valStream <- v:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func gen(nums ...int) chan int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, n := range nums {
			out <- n
			// you can comment the below
			// to check if or-done channel works
			time.Sleep(50 * time.Millisecond)
		}
	}()
	return out
}
