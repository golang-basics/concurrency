package main

import (
	"fmt"
	"time"
)

func work() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("printing stuff")
}

func main() {
	go work()
	time.Sleep(100 * time.Millisecond)
	fmt.Println("done waiting, main exits")
}
