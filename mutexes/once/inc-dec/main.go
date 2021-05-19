package main

import (
	"fmt"
	"sync"
)

func main() {
	var i int
	var once sync.Once
	inc := func() {
		i++
	}
	dec := func() {
		i--
	}

	once.Do(inc)
	once.Do(dec)

	fmt.Println("i:", i)
}
