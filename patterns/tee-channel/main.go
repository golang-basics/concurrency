// The tee channel pattern works exactly like the Linux tee command.
// The tee command writes the output to STDOUT and a list of specified FILES.
// Respectively the tee channel will take an input channel and clone that into a
// specified number of channel for further pipeline operations.

package main

import (
	"fmt"
	"sync"

	"concurrency/patterns/pipeline"
)

func main() {
	done := make(chan struct{})
	defer close(done)

	nums := pipeline.Gen(1, 2, 3, 4)
	out := tee(done, nums, 3)

	for v := range pipeline.Inc(out[0]) {
		fmt.Println("incremented:", v)
	}
	for v := range pipeline.Dec(out[1]) {
		fmt.Println("decremented:", v)
	}
	for v := range pipeline.Sq(out[2]) {
		fmt.Println("squared:", v)
	}
}

// maybe generate n clone channels
// for example: ls | tee f1.txt f2.txt f3.txt
func tee(done <-chan struct{}, in <-chan int, n int) []chan int {
	var wg sync.WaitGroup
	wg.Add(1)
	values := make([]int, 0)
	go func() {
		defer wg.Done()
		for val := range in {
			select {
			case <-done:
				return
			default:
				values = append(values, val)
			}
		}
	}()
	wg.Wait()

	channels := make([]chan int, 0)
	defer func() {
		for _, ch := range channels {
			close(ch)
		}
	}()

	for i := 0; i < n; i++ {
		channels = append(channels, make(chan int, len(values)))
	}
	for _, val := range values {
		select {
		case <-done:
			break
		default:
			for _, ch := range channels {
				ch <- val
			}
		}
	}
	return channels
}
