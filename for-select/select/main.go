package main

import "fmt"

func main() {
	ch := make(chan int)
	go write(ch)

	select {
	case <-ch:
		fmt.Println("wrote to ch")
	}
}

func write(ch chan int) {
	ch <- 1
}
