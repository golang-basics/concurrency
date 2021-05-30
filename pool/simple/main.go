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

// Here are couple of sync.Pool use cases
// 1. You won't notice performance improvements on simple data (faster using normal allocations)
// 2. You won't notice performance improvements (actually slower) if GC is triggered often
// 3. You must have some sort of reset function to clear dirty object data fetched from Pool
// 4. You usually get a lot of performance benefits when running things in parallel and when
// dealing with relatively big objects which otherwise have a big cost allocating/deallocating frequently
func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("creating new object")
			return struct{}{}
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
