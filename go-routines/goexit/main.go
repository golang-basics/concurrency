package main

import (
	"fmt"
	"runtime"
	"sync"
)

// calling runtime.Goexit inside main will crash the program if
// no go routines are running, resulting in a deadlock
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer fmt.Println("go routine exited")
		var i int
		for {
			if i == 100 {
				runtime.Goexit()
				return
			}
			i++
		}
	}()
	wg.Wait()
}
