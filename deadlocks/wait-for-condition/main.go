package main

import "fmt"

func main() {
	value, dep := make(chan int), make(chan struct{})
	go func() {
		<-dep
		// value never gets produced because the above blocks forever
		value <- 1
	}()
	// in order to execute fmt.Println statement we need value
	// the condition is never met, and we have a deadlock
	fmt.Println(1 + <-value)
}
