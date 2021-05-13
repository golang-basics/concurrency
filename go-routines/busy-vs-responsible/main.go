package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	c := make(chan int)
	//go busy(c)
	go responsible(c)
	c <- 1
	close(c)
	time.Sleep(time.Second)
	fmt.Println("number of go routines", runtime.NumGoroutine())
}

func busy(in chan int) {
	defer fmt.Println("I am never done")
	for {
		select {
		case v := <-in:
			fmt.Println("value:", v)
		}
	}
}

func responsible(in chan int) {
	defer fmt.Println("I am responsibly done")
	for {
		select {
		case v, ok := <-in:
			if !ok {
				return
			}
			fmt.Println("value:", v)
		}
	}
}
