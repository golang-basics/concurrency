package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	john := person{name: "John"}
	alice := person{name: "Alice"}
	hallway(john, alice)
}

type person struct {
	name    string
	cadence *sync.Cond
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
	atomic.AddInt32(direction, 1)
	p.wait()
	// nobody seems to have walked this direction, I can take it
	if atomic.LoadInt32(direction) == 1 {
		return true
	}
	p.wait()
	// let me step back, someone else is trying to walk the same direction
	atomic.AddInt32(direction, -1)
	return false
}

func (p person) wait() {
	p.cadence.L.Lock()
	p.cadence.Wait()
	p.cadence.L.Unlock()
}

func hallway(people ...person) {
	done := make(chan struct{})
	defer close(done)
	cadence := sync.NewCond(&sync.Mutex{})
	go func() {
		for {
			select {
			case <-done:
				return
			case <-time.Tick(500 * time.Millisecond):
				cadence.Broadcast()
			}
		}
	}()

	var wg sync.WaitGroup
	var left, right int32
	for _, p := range people {
		wg.Add(1)
		p.cadence = cadence
		go func(p person) {
			p.walk(&left, &right)
			wg.Done()
		}(p)
	}
	wg.Wait()
}
