package main

import (
	"fmt"
	"time"
)

func main() {
	queue := make(chan string, 2)
	people := []string{"John", "Jane", "Mike", "Steve", "Alex"}

	// people come and wait to be served
	go func() {
		defer close(queue)
		for _, p := range people {
			time.Sleep(time.Second)
			fmt.Println(p, "starts waiting")
			queue <- p
		}
	}()

	// the restaurant is trying to serve people waiting in the queue
	// the maximum capacity of serving people is 2
	// the restaurant serves people slower, then they arrive in the queue
	for person := range queue {
		time.Sleep(3 * time.Second)
		fmt.Println("serving:", person)
	}
	time.Sleep(time.Second)
	fmt.Println("closing the restaurant")
}
