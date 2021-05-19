package main

import (
	"fmt"
	"runtime"
	"sync"
)

func main() {
	runtime.GOMAXPROCS(1)
	p := &sync.Pool{
		New: func() interface{} {
			fmt.Println("create object")
			return struct{}{}
		},
	}
	// object get created here
	obj := p.Get()
	p.Put(obj)

	// object gets reused here
	obj = p.Get()
	p.Put(obj)

	runtime.GC()
	fmt.Println("garbage collecting")
	// object gets reused here
	obj = p.Get()
	p.Put(obj)
}
