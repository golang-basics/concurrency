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
	inc, dec, sq := make([]int, 0), make([]int, 0), make([]int, 0)

	var wg sync.WaitGroup
	wg.Add(3)
	go func() {
		defer wg.Done()
		for v := range pipeline.Inc(done, ch1) {
			// run it inside go routine if order does not matter
			// fmt.Println("incremented:", v)
			inc = append(inc, v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range pipeline.Dec(done, ch2) {
			// run it inside go routine if order does not matter
			// fmt.Println("decremented:", v)
			dec = append(dec, v)
		}
	}()
	go func() {
		defer wg.Done()
		for v := range pipeline.Sq(done, ch3) {
			// run it inside go routine if order does not matter
			// fmt.Println("squared:", v)
			sq = append(sq, v)
		}
	}()
	wg.Wait()

	// We stored the values inside slices
	// to preserve the order after all channels are drained.
	// If order is not important, the processing could be
	// done inside the corresponding go routine
	fmt.Println("incremented:")
	for _, v := range inc {
		fmt.Println(v)
	}
	fmt.Println("decremented:")
	for _, v := range dec {
		fmt.Println(v)
	}
	fmt.Println("squared:")
	for _, v := range sq {
		fmt.Println(v)
	}
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
		// we could use a batching approach instead
		// but it would not be as performant, hence streams
		// take advantage of speed and memory
		// the bigger the incoming stream the worse would've
		// been the performance, if we used batching
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
