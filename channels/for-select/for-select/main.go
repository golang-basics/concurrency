package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan int)
	go write(ch)

	timeout := time.NewTimer(3 * time.Second)
	ticker := time.NewTicker(500 * time.Millisecond)
	for {
		select {
		case at := <-ticker.C:
			fmt.Println("tick", at)
		case <-timeout.C:
			fmt.Println("time is up")
			return
		case v := <-ch:
			fmt.Println("wrote:", v)
		}
	}
}

func write(ch chan int) {
	for i := 0; i < 100; i++ {
		ch <- i
		time.Sleep(200 * time.Millisecond)
	}
}
