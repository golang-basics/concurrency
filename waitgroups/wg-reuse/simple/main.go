package main

import (
	"sync"
	"time"
)

// This is a small extras from the official docs
// --------------------------------------------
// Note that calls with a positive delta that occur when the counter is zero
// must happen before a Wait. Calls with a negative delta, or calls with a
// positive delta that start when the counter is greater than zero, may happen
// at any time.
func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		time.Sleep(time.Millisecond)
		wg.Add(-1)
		wg.Add(1)
	}()
	wg.Wait()
}
