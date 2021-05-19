package main

import (
	"fmt"
	"sync"
)

// sync.Pool is a concurrent safe type
// so feel free to run this program with -race flag
// go run -race main.go

// Also do not make any false assumptions about the Get relation to Set
// to break it to you here are some comment from the source code
// -------------------------------------------------------------
// Get may choose to ignore the pool and treat it as empty.
// Callers should not assume any relation between values passed to Put and
// the values returned by Get.
func main() {
	var wg sync.WaitGroup
	p := sync.Pool{}
	p.New = func() interface{} {
		return "nothing"
	}

	wg.Add(3)
	go func() {
		p.Put("object 1")
		wg.Done()
	}()
	go func() {
		p.Put("object 2")
		wg.Done()
	}()
	go func() {
		p.Put("object 3")
		wg.Done()
	}()
	wg.Wait()

	wg.Add(3)
	go func() {
		fmt.Println(p.Get())
		wg.Done()
	}()
	go func() {
		fmt.Println(p.Get())
		wg.Done()
	}()
	go func() {
		fmt.Println(p.Get())
		wg.Done()
	}()
	wg.Wait()

	fmt.Println(p.Get())
}
