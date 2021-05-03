package main

import (
	"fmt"
	"time"
)

func main() {
	buffered(2) // go routines will not block
	buffered(0) // go routines will block
	time.Sleep(time.Second)
	fmt.Println("DONE")
}

func buffered(n int) {
	c := make(chan int, n)
	go func() {
		c <- 1
		fmt.Println("go routine 1")
		// will return immediately if chan is buffered
	}()
	go func() {
		c <- 2
		fmt.Println("go routine 2")
		// will return immediately if chan is buffered
	}()
	time.Sleep(time.Second)
	fmt.Println(<-c)
	time.Sleep(time.Second)
	fmt.Println(<-c)
}
