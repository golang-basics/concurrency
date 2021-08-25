package main

import "sync"

// sync.Cond uses a mutex, so it pretty much has the exact same mutex issues
func main() {
	cond := sync.NewCond(new(sync.Mutex))
	cond.L.Lock()
	cond.L.Lock()
}
