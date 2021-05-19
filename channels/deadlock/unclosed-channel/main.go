package main

import "fmt"

func main() {
	ch := make(chan int)
	for v := range ch {
		fmt.Println(v)
	}
}

func write(ch chan int) {
	ch <- 1
	ch <- 2
	ch <- 3
	// close(ch) must be called
}
