package main

import (
	"fmt"
	"time"
)

func main() {
	go work() // fork point
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done waiting, main exits")
	// join point
}

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing some stuff")
}
