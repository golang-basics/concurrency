package main

import (
	"fmt"
	"sync"
)

// Here are some important things to keep in mind about Broadcast
// It must be called only after we made sure all the go routines
// are ready for the rendezvous point. Calling Broadcast before
// the rendezvous point is met in conjunction with a wait group
// will result in a deadlock, because Broadcast will  only notify
// go routines that area ready, which may leave some go routines
// which have calls to wg.Done().
// Which means calls to wg.Done() never happen,
// which results in wg.Wait() call waiting forever, thus resulting in a deadlock
// In the below example we used 2 wait groups
// 1. For making sure all go routines are ready and waiting for the rendezvous
// 2. For making sure we wait till all workers / go routines execute before
// exiting the main function
func main() {
	cond := sync.NewCond(&sync.Mutex{})
	var wg sync.WaitGroup

	wg.Add(2)
	register(cond, func() {
		defer wg.Done()
		fmt.Println("worker 1")
	})
	register(cond, func() {
		defer wg.Done()
		fmt.Println("worker 2")
	})
	cond.Broadcast()

	wg.Wait()
}

func register(cond *sync.Cond, fn func()) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		wg.Done()
		cond.L.Lock()
		cond.Wait()
		fn()
		cond.L.Unlock()
	}()
	wg.Wait()
}
