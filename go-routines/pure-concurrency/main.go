package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

var wg sync.WaitGroup

func main() {
	runtime.GOMAXPROCS(1)

	wg.Add(3)
	now := time.Now()
	go func() {
		task1()
		wg.Done()
	}()
	go func() {
		task2()
		wg.Done()
	}()
	go func() {
		task3()
		wg.Done()
	}()
	wg.Wait()
	fmt.Println("elapsed", time.Now().Sub(now))

	fmt.Printf("\n\n")

	now = time.Now()
	task1()
	task2()
	task3()
	fmt.Println("elapsed", time.Now().Sub(now))
}

func task1() {
	fmt.Println("before task 1")
	time.Sleep(1 * time.Second)
	fmt.Println("after task 1")
}

func task2() {
	fmt.Println("before task 2")
	time.Sleep(3 * time.Second)
	fmt.Println("after task 2")
}

func task3() {
	fmt.Println("before task 3")
	time.Sleep(2 * time.Second)
	fmt.Println("after task 3")
}
