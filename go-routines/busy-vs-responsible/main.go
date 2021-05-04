package main

import (
	"fmt"
	"time"
)

func main() {
	c := make(chan int)
	//go busy(c)
	go responsible(c)
	c <- 1
	close(c)
	time.Sleep(1 * time.Second)
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
