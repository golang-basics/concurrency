// The bridge channel pattern is pretty self explanatory. It acts a channel bridge
// for a bunch of other channels. This pattern removes the effort of intermediate
// individual channel reading, bridging everything from a given amount of channels
// to a single easily consumable channel.
// Working with a channel of channels is most of the times cumbersome,
// and the bridge channel pattern lets us focus only on what's needed
// as opposed to wasting time on destructuring the channel of channels

package main

import (
	"fmt"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	for v := range bridge(done, genStreams()) {
		fmt.Println("value:", v)
	}
}

func bridge(done <-chan struct{}, chanStream <-chan <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			var stream <-chan int
			select {
			case maybeStream, ok := <-chanStream:
				if ok == false {
					return
				}
				stream = maybeStream
			case <-done:
				return
			}
			for val := range stream {
				select {
				case out <- val:
				case <-done:
				}
			}
		}
	}()
	return out
}

// generate a channel of 10 channels each streaming 3 values
func genStreams() <-chan <-chan int {
	out := make(chan (<-chan int))
	go func() {
		defer close(out)
		for i := 1; i <= 10; i++ {
			stream := make(chan int, 3)
			stream <- i
			stream <- i + 1
			stream <- i + 2
			close(stream)
			out <- stream
		}
	}()
	return out
}
