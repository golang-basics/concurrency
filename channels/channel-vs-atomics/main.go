package main

import (
	"fmt"
)

type set struct {
	value int32
	done  chan struct{}
}

type get struct {
	value chan int32
}

// try running the example like this: go run -race main.go
func main() {
	var value int32
	setChan := make(chan set)
	getChan := make(chan get)

	go func() {
		for {
			select {
			case op := <-getChan:
				op.value <- value
			case op := <-setChan:
				value = op.value
				op.done <- struct{}{}
			}
		}
	}()
	for i := 0; i < 1000; i++ {
		go func(i int32) {
			op := set{
				value: i,
				done:  make(chan struct{}),
			}
			setChan <- op
			<-op.done
		}(int32(i))
	}

	op := get{value: make(chan int32)}
	getChan <- op
	fmt.Println("value:", <-op.value)
}
