package main

import (
	"fmt"
	"sync"
)
// A sync.Pool may be a misunderstood data type and not clear about its purpose.
// Here a small excerpt from the official docs comments:
// ----------------------------------------------------
// Pool's purpose is to cache allocated but unused items for later reuse,
// relieving pressure on the garbage collector. That is, it makes it easy to
// build efficient, thread-safe free lists. However, it is not suitable for all
// free lists.
func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new object")
			return struct {}{}
		},
	}
	// this will invoke New
	pool.Get()

	// this will invoke New
	obj := pool.Get()
	pool.Put(obj)

	// this will not invoke the New func
	pool.Get()
}
