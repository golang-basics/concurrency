package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

func init() {
	// Usually the call to LockOSThread is made here in the init function
	// most of the times when using C code, which manages its own concurrency
	// to avoid having the main go routine scheduled on different threads
	// we lock it on a specific thread so that calls to CGO and main go routine
	// execute on the same thread.
	//runtime.LockOSThread()
}

func main() {
	work := make(chan struct{})
	done := make(chan struct{})
	defer close(done)
	go worker(work, done)

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go producer(work, &wg)
	}
	wg.Wait()
}

func worker(work, done chan struct{}) {
	// To avoid a lot of context switching
	// sometimes it could be useful to run some heavy work
	// on a thread that is not interrupted or context switched
	// by locking the OS thread.
	// Locking the OS thread does not always mean it will execute on the same
	// Processor (CORE), the Scheduler may switch it between multiple CORES.
	runtime.LockOSThread()
	for {
		select {
		case <-done:
			runtime.UnlockOSThread()
			return
		case <-work:
			time.Sleep(500 * time.Millisecond)
			fmt.Println("doing some heavy work")
		}
	}
}

func producer(work chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	work <- struct{}{}
}
