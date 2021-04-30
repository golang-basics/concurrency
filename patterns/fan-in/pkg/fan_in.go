// The FAN-IN pattern states that a function receives multiple channels as inputs
// It reads each input and sends all the values into 1 final output channel

package pkg

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
