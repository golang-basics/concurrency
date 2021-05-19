package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go read(ch)
	ch <- 1
	ch <- 2
	ch <- 3
	time.Sleep(time.Second)
}

func read(ch <-chan int) {
	for v := range ch {
		fmt.Println(v)
		// ch<-1 it result in a compilation error: send to receive only chan
	}
}
