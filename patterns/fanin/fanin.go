package fanin

import (
	"sync"
)

// FanIn reads from multiple channels and writes into 1 final channel
// The FAN-IN aka Multiplexing pattern states that a function receives multiple channels as inputs
// It reads each input and sends all the values into 1 final output channel
func FanIn(done chan struct{}, inputs ...<-chan int) <-chan int {
	out := make(chan int)
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(numbers <-chan int) {
			defer wg.Done()
			for n := range numbers {
				select {
				case <-done:
					return
				case out <- n:
				}
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}
