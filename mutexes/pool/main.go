package main

import (
	"fmt"
	"sync"
)

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
