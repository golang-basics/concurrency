package pkg

// FanIn reads from multiple channels and writes into 1 final channel
// The FAN-IN aka Multiplexing pattern states that a function receives multiple channels as inputs
// It reads each input and sends all the values into 1 final output channel
func FanIn(inputs ...<-chan int) <-chan int {
	out := make(chan int)
	done := make(chan struct{})
	for _, in := range inputs {
		go func(numbers <-chan int) {
			for n := range numbers {
				out <- n
			}
			done <- struct{}{}
		}(in)
	}
	go func() {
		for i := 0; i < len(inputs); i++ {
			<-done
		}
		close(done)
		close(out)
	}()
	return out
}
