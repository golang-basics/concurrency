package main

import "sync"

// sync.Cond uses a mutex, so it pretty much has the exact same mutex issues
func main() {
	cond := sync.NewCond(new(sync.Mutex))
	cond.L.Lock()
	// the call to Wait() does 3 things
	// 1. Call Unlock() on Cond locker
	// 2. notify the wait list
	// 3. Call Lock() on Cond locker
	cond.Wait()
}
