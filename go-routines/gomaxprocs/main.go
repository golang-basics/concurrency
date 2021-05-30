package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)
	go task1()
	go task2()
	go task3()
	time.Sleep(time.Second)
	//for{}
}

func task1() {
	fmt.Println("task 1")
}

func task2() {
	fmt.Println("task 2")
}

func task3() {
	fmt.Println("task 3")
}
