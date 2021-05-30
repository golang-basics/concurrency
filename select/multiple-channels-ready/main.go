package main

import "fmt"

func main() {
	count1, count2 := 0, 0
	stream1, stream2 := make(chan struct{}), make(chan struct{})
	close(stream1)
	close(stream2)

	for i := 0; i < 1000; i++ {
		select {
		case <-stream1:
			count1++
		case <-stream2:
			count2++
		}
	}

	fmt.Println("count1:", count1)
	fmt.Println("count2:", count2)
}
