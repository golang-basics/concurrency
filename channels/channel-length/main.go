package main

import "fmt"

func main() {
	c := make(chan int, 100)
	for i := 0; i < 34; i++ {
		c <- 0
	}
	fmt.Println(len(c))
}
