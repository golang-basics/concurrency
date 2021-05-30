package main

import "fmt"

func main() {
	ch := make(chan int)
	go write(ch)
	// range expects the producer
	// to close the outbound channel
	// otherwise it will result in deadlock
	// Note: the for loop does not
	// need an exit condition for channels
	// it just exits when the channel is closed
	for i := range ch {
		fmt.Println(i)
	}
}

func write(ch chan int) {
	for i := 0; i < 3; i++ {
		ch <- i
	}
	close(ch)
}
