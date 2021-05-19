package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int, 1)
	go write(ch)

	time.Sleep(2 * time.Second)
	for v := range ch {
		time.Sleep(1 * time.Second)
		// channel read/write executes faster than writing to STD OUT
		fmt.Println("read:", v)
	}
}

func write(ch chan int) {
	for i := 1; i < 4; i++ {
		fmt.Println("writing:", i)
		fmt.Println("---")
		ch <- i
	}
	close(ch)
}
