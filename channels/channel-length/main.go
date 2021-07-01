package main

import "fmt"

func main() {
	c := make(chan int, 100)
	for i := 0; i < 20; i++ {
		c <- i
	}
	fmt.Println("channel capacity", cap(c))
	fmt.Println("channel length", len(c))
}
