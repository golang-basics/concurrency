package main

import (
	"fmt"
	"sync"
)

type person struct {
	name string
}

func main() {
	pool := &sync.Pool{
		New: func() interface{} {
			return nil
		},
	}

	pool.Put(&person{name: "John"})
	pool.Put(&person{name: "Amy"})
	pool.Put(&person{name: "Steve"})

	for {
		p, ok := pool.Get().(*person)
		if !ok {
			fmt.Println("we're closed")
			return
		}
		fmt.Println("serving:", p.name)
	}
}
