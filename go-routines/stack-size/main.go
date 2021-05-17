package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	var c <-chan struct{}
	noop := func() {
		wg.Done()
		<-c
	}
	const numGoRoutines = 1e4
	wg.Add(numGoRoutines)
	before := stackSize()
	for i := 0; i < numGoRoutines; i++ {
		go noop()
	}
	wg.Wait()
	after := stackSize()
	fmt.Printf("%.3fkb", float64(after-before)/numGoRoutines/1024)
}

func stackSize() uint64 {
	runtime.GC()
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return m.StackSys
}
