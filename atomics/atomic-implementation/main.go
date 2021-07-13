package main

import (
	"fmt"
	"sync"
)

// taken from the source code
// LoadInt64
// https://github.com/golang/go/blob/master/src/sync/atomic/doc.go#L114
// https://github.com/golang/go/blob/master/src/sync/atomic/asm.s#L61
// https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.go#L28

// StoreInt64
// https://github.com/golang/go/blob/master/src/sync/atomic/doc.go#L132
// https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.go#L101
// https://github.com/golang/go/blob/master/src/runtime/internal/atomic/atomic_amd64.s#L171
func StoreInt64(addr *int64, value int64)

// to test this example, make sure to build it using the -race flag
// go build -race -o exec
// ./exec
func main() {
	var count int64
	var wg sync.WaitGroup

	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(i int) {
			defer wg.Done()
			StoreInt64(&count, int64(i+1))
		}(i)
	}

	wg.Wait()
	fmt.Println("count", count)
}
