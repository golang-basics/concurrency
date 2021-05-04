package main

import "fmt"

func main() {
	c := make(chan int)
	close(c)
	// this gets ignored
	<-c
	// this gets ignored as well
	for _ = range c {
		fmt.Println("I won't print")
	}
	// this results in an error
	c <- 1
}
