package main

import "sync"

// Calling Wait() requires the Mutex to be locked,
// Not doing so results in a panic.
// In fact Wait() calls Unlock() in the first place
func main() {
	cond := sync.NewCond(new(sync.Mutex))
	cond.Wait()
	//cond.L.Unlock()
}
