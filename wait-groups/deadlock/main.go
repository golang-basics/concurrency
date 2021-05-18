package main

import (
	"sync"
)

// Calling Done more times than Add panics
// Calling Done less times than Add results in deadlock
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	wg.Wait()
}
