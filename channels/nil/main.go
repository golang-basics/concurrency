package main

import "fmt"

func main() {
	var c chan int
	// results in deadlock (blocks forever)
	<-c
	// results in deadlock (blocks forever)
	for _ = range c {
		fmt.Println("I won't print")
	}
	// results in deadlock (blocks forever)
	c <- 1
	// panics: close on nil channel
	close(c)
}
