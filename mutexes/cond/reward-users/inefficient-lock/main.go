package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var mu sync.Mutex
	var wg sync.WaitGroup
	var checks int32
	users := make([]string, 0)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			mu.Lock()
			time.Sleep(5 * time.Millisecond)
			users = append(users, fmt.Sprintf("user %d", i+1))
			mu.Unlock()
		}
	}()
	go func() {
		defer wg.Done()
		// burning unnecessary CPU cycles
		for {
			mu.Lock()
			atomic.AddInt32(&checks, 1)
			if len(users) >= 100 {
				reward(users[:10])
				mu.Unlock()
				return
			}
			mu.Unlock()
		}
	}()

	wg.Wait()
	fmt.Println("number of checks", checks)
}

func reward(users []string) {
	for _, user := range users {
		fmt.Println("rewarding:", user)
	}
}
