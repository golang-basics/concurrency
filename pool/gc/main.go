package main

import (
	"fmt"
	"runtime"
	"sync"
)

// Pools are always cleared before Garbage Collector kicks in
// Have a look inside the source code
// https://github.com/golang/go/blob/master/src/runtime/mgc.go#L1547
// https://github.com/golang/go/blob/master/src/sync/pool.go#L233
func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create object")
			return struct{}{}
		},
	}

	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		runtime.GC()
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println("go routine", i)
			obj := pool.Get()
			pool.Put(obj)
		}(i)
	}
	wg.Wait()

	runtime.GC()
	fmt.Println("after gc")
	obj := pool.Get()
	pool.Put(obj)
}
