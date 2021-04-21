package main

import (
	"fmt"
	"time"
)

func main() {
	done := make(chan struct{})
	go task1(done)
	<-done
	go task2(done)
	<-done
	go task3(done)
	<-done
}

func task1(done chan struct{}) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1")
	done <- struct{}{}
}

func task2(done chan struct{}) {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("task 2")
	done <- struct{}{}
}

func task3(done chan struct{}) {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("task 3")
	done <- struct{}{}
}
