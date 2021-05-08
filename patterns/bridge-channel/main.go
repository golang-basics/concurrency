// The bridge channel pattern is pretty self explanatory. It acts a channel bridge
// for a bunch of other channels. This pattern removes the effort of intermediate
// individual channel reading, bridging everything from a given amount of channels
// to a single easily consumable channel.
// Working with a channel of channels is most of the times cumbersome,
// and the bridge channel pattern lets us focus only on what's needed

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
	valStream := make(chan int)
	go func() {
		defer close(valStream)
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
				case valStream <- val:
				case <-done:
				}
			}
		}
	}()
	return valStream
}

func genStreams() <-chan <-chan int {
	chanStream := make(chan (<-chan int))
	go func() {
		defer close(chanStream)
		for i := 1; i <= 10; i++ {
			stream := make(chan int, 3)
			stream <- i
			stream <- i + 1
			stream <- i + 2
			close(stream)
			chanStream <- stream
		}
	}()
	return chanStream
}
