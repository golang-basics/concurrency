package main

import (
	"fmt"
	"sync"
)

// Taken from sync/Once.Do source code
// Because no call to Do returns until the one call to f returns, if f causes
// Do to be called, it will deadlock.
func main() {
	//deadlock1()
	deadlock2()
}

func deadlock1() {
	var once sync.Once
	once.Do(func() {
		// This Do won't return because the next Do is called,
		// thus causing the main func to deadlock
		fmt.Println("executing func once")
		once.Do(func() {
			fmt.Println("will it work?")
		})
	})
}

func deadlock2() {
	var onceA, onceB sync.Once
	var b func()
	a := func() {
		fmt.Println("before a")
		onceB.Do(b)
		fmt.Println("after a")
	}
	b = func() {
		fmt.Println("before b")
		// onceA.Do was already executed once
		// thus this call will never return
		// causing a deadlock in main
		onceA.Do(a)
		fmt.Println("after b")
	}

	onceA.Do(a)
}
