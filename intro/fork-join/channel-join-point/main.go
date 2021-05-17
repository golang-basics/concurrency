package main

import (
	"fmt"
	"time"
)

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing stuff")
}

func main() {
	done := make(chan struct{})
	go func() {
		work()
		done <- struct{}{}
	}()
	<-done
	fmt.Println("done waiting, main exits")
}
