package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

// Taken from the standard library source code
// ---------------------------------------------------------------
// Because c.L is not locked when Wait first resumes, the caller
// typically cannot assume that the condition is true when
// Wait returns. Instead, the caller should Wait in a loop:
//
//    c.L.Lock()
//    for !condition() {
//        c.Wait()
//    }
//    ... make use of condition ...
//    c.L.Unlock()
//
func main() {
	var wg sync.WaitGroup
	var checks int32
	cond := sync.NewCond(new(sync.Mutex))
	users := make([]string, 0)

	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 0; i < 100; i++ {
			cond.L.Lock()
			time.Sleep(5 * time.Millisecond)
			users = append(users, fmt.Sprintf("user %d", i+1))
			cond.L.Unlock()
			cond.Signal()
		}
	}()
	go func() {
		defer wg.Done()
		cond.L.Lock()
		for len(users) < 100 {
			atomic.AddInt32(&checks, 1)
			cond.Wait()
		}
		reward(users[:10])
		cond.L.Unlock()
	}()
	wg.Wait()
	fmt.Println("number of checks", checks)
}

func reward(users []string) {
	for _, user := range users {
		fmt.Println("rewarding:", user)
	}
}
