package generators

// IntRepeat generates an infinite number of repeated values till stopped via the done channel
func IntRepeat(done <-chan struct{}, values ...int) <-chan int {
	valueStream := make(chan int)
	go func() {
		defer close(valueStream)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case valueStream <- v:
				}
			}
		}
	}()
	return valueStream
}

// IntTake extracts a sub stream of values out of the value stream till it reaches a certain amount or stopped
func IntTake(done <-chan struct{}, valueStream <-chan int, num int) <-chan int {
	takeStream := make(chan int)
	go func() {
		defer close(takeStream)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case takeStream <- <-valueStream:
			}
		}
	}()
	return takeStream
}
