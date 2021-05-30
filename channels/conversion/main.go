package main

func main() {
	var receiveChan <-chan int
	var sendChan chan<- int
	intChan := make(chan int)

	// 2 directional channel gets converted into 1 directional channel
	receiveChan = intChan
	sendChan = intChan

	// to avoid compilation errors
	receiveChan = receiveChan
	sendChan = sendChan
}
