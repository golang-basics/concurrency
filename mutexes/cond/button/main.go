package main

import (
	"fmt"
	"sync"
)

type button struct {
	clicked *sync.Cond
}

func main() {
	btn := button{clicked: sync.NewCond(&sync.Mutex{})}
	var wg sync.WaitGroup

	wg.Add(3)
	subscribe(btn.clicked, handler1(&wg))
	subscribe(btn.clicked, handler2(&wg))
	subscribe(btn.clicked, handler3(&wg))

	btn.clicked.Broadcast()
	wg.Wait()
}

func subscribe(cond *sync.Cond, fn func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		cond.L.Lock()
		defer cond.L.Unlock()
		cond.Wait()
		fn()
	}()
	wg.Wait()
}

func handler1(wg *sync.WaitGroup) func() {
	return func() {
		fmt.Println("closing popup window")
		wg.Done()
	}
}

func handler2(wg *sync.WaitGroup) func() {
	return func() {
		fmt.Println("redirecting to a different page")
		wg.Done()
	}
}

func handler3(wg *sync.WaitGroup) func() {
	return func() {
		fmt.Println("display available options")
		wg.Done()
	}
}
