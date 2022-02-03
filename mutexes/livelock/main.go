package main

import (
	"fmt"
	"sync"
	"time"
)

// Also try running this example with the -race flag
// go run -race main.go
func main() {
	john := person{name: "John"}
	alice := person{name: "Alice"}
	hallway(john, alice)
}

func hallway(people ...person) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var left, right int32
	for _, p := range people {
		wg.Add(1)
		p.mu = &mu
		go func(p person) {
			defer wg.Done()
			p.walk(&left, &right)
		}(p)
	}
	wg.Wait()
}

type person struct {
	name string
	mu   *sync.Mutex
}

func (p *person) walk(left, right *int32) {
	for i := 0; i < 5; i++ {
		if p.step("left", left) || p.step("right", right) {
			fmt.Println(p.name, "successfully walked")
			return
		}
	}
	fmt.Println(p.name, "gives up trying")
}

func (p *person) step(directionName string, direction *int32) bool {
	fmt.Println(p.name, "trying to walk", directionName)
	// let me walk this direction
	p.advance(direction, 1)

	// nobody seems to have walked this direction, I can take it
	var d int32
	p.mu.Lock()
	d = *direction
	p.mu.Unlock()
	if d == 1 {
		return true
	}
	time.Sleep(100 * time.Millisecond)

	// let me step back, someone else is trying to walk the same direction
	p.advance(direction, -1)
	return false
}

func (p *person) advance(direction *int32, increment int32) {
	p.mu.Lock()
	*direction += increment
	p.mu.Unlock()
	time.Sleep(500 * time.Millisecond)
}
