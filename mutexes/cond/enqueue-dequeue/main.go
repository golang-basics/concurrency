package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	cond := sync.NewCond(&sync.Mutex{})
	queue := make([]int, 0, 10)
	dequeue := func(delay time.Duration) {
		time.Sleep(delay)
		cond.L.Lock()
		fmt.Println("dequeued", queue[0])
		queue = queue[1:]
		cond.L.Unlock()
		cond.Signal()
	}

	for i := 0; i < 10; i++ {
		cond.L.Lock()
		// wait for dequeue if queue is full => 2
		for len(queue) == 2 {
			cond.Wait()
		}
		fmt.Println("enqueue", i)
		queue = append(queue, i)
		go dequeue(time.Second)
		cond.L.Unlock()
	}
}
