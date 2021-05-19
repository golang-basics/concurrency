package main

import (
	"fmt"
)

func main() {
	c := make(chan bool)
	go task1(c)
	<-c
	go task2(c)
	<-c
}

func task1(c chan bool) {
	fmt.Println("task 1")
	c <- true
}

func task2(c chan bool) {
	fmt.Println("task 2")
	c <- false
}
