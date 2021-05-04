package main

import (
	"fmt"
	"time"
)

func main() {
	go worker()
	go func() {
		fmt.Println("print in anonymous function")
	}()
	time.Sleep(time.Second)
}

func worker() {
	fmt.Println("print in worker function")
}
