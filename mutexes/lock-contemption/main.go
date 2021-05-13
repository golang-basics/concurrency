package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	var mu sync.Mutex

	wg.Add(2)
	go func() {
		mu.Lock()
		time.Sleep(time.Second)
		fmt.Println("go routine 1 releasing lock:", time.Now())
		mu.Unlock()
		wg.Done()
	}()
	go func() {
		fmt.Println("go routine 2 acquiring lock:", time.Now())
		mu.Lock()
		fmt.Println("go routine 2 done:", time.Now())
		mu.Unlock()
		wg.Done()
	}()

	wg.Wait()
}
