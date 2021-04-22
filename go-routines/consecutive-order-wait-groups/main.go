package main

import (
	"fmt"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	wg.Add(1)
	go task1()
	wg.Wait()

	wg.Add(1)
	go task2()
	wg.Wait()

	wg.Add(1)
	go task3()
	wg.Wait()
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1")
	wg.Done()
}

func task2() {
	time.Sleep(50 * time.Millisecond)
	fmt.Println("task 2")
	wg.Done()
}

func task3() {
	time.Sleep(10 * time.Millisecond)
	fmt.Println("task 3")
	wg.Done()
}
