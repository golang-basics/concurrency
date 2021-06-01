package main

import (
	"fmt"
	"time"
)

func main() {
	go longTask()
	immediateTask()
	time.Sleep(time.Second)
}

func immediateTask() {
	fmt.Println("I executed immediately")
}

func longTask() {
	time.Sleep(500 * time.Millisecond)
	fmt.Println("I executed after 500ms")
}
