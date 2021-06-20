package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for {
		// this causes the wait group to call Wait before all
		// calls to Done are finished, thus resulting in an error
		go work(&wg)
	}
}

func work(wg *sync.WaitGroup) {
	wg.Add(3)
	go task1(wg)
	go task2(wg)
	go task3(wg)
	wg.Wait()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	return
}
func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	return
}
func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	return
}
