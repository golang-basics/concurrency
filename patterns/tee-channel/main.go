// The tee channel pattern works exactly like the Linux tee command.
// The tee command writes the output to STDOUT and a list of specified FILES.
// Respectively the tee channel will take an input channel and clone that into a
// specified number of channels for further pipeline operations.

package main

import (
	"fmt"
	"sync"

	"concurrency/patterns/pipeline"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	nums := pipeline.Gen(done, 1, 2, 3, 4)
	out := tee(done, nums, 3)
	ch1, ch2, ch3 := out[0], out[1], out[2]

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for v := range pipeline.Inc(done, ch1) {
			fmt.Println("incremented:", v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range pipeline.Dec(done, ch2) {
			fmt.Println("decremented:", v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range pipeline.Sq(done, ch3) {
			fmt.Println("squared:", v)
		}
	}()
	wg.Wait()
}

// The tee channel works pretty much like the tee Unix command.
// Try this command: ls | tee f1.txt f2.txt f3.txt
// which will list the current directory contents to stdout and the above files
func tee(done <-chan struct{}, in <-chan int, n int) []chan int {
	channels := make([]chan int, 0)
	for i := 0; i < n; i++ {
		channels = append(channels, make(chan int))
	}
	go func() {
		// close all outbound channels at the end
		defer func() {
			for _, in := range channels {
				close(in)
			}
		}()

		// range over the streamed values
		for val := range in {
			select {
			case <-done:
				return
			default:
				for _, ch := range channels {
					// Keep in mind, each channel must have a go routine
					// that reads from it, any kind of blocking operation/ignorance
					// will result in a deadlock. So make sure to use all channels correctly.
					ch <- val
				}
			}
		}
	}()
	return channels
}
