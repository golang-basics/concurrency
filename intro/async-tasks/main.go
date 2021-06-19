package main

import (
	"fmt"
	"time"
)

func main() {
	//asyncNotWorking()
	asyncTmpFix()
}

func asyncNotWorking() {
	now := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()
	// WHY THE CODE ABOVE DOES NOT WORK?
	// WHAT COULD POSSIBLY GO WRONG?
	fmt.Println("elapsed", time.Now().Sub(now))
}

func asyncTmpFix() {
	now := time.Now()
	go task1()
	go task2()
	go task3()
	go task4()
	// ALSO WHAT THE HECK IS WRONG with the order
	fmt.Println("elapsed", time.Now().Sub(now))
	// WHY? FOR WHAT?
	time.Sleep(time.Second)
}

func task1() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 1")
}

func task2() {
	time.Sleep(200 * time.Millisecond)
	fmt.Println("task 2")
}

func task3() {
	fmt.Println("task 3")
}

func task4() {
	time.Sleep(100 * time.Millisecond)
	fmt.Println("task 4")
}
