package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	//mu := basicMutex{}
	mu := rwMutex{}
	mu.store(10)

	wg.Add(3)
	before := time.Now()
	go func() {
		defer wg.Done()
		value := mu.load()
		fmt.Println("value go routine 1:", value)
	}()
	go func() {
		defer wg.Done()
		value := mu.load()
		fmt.Println("value go routine 2:", value)
	}()
	go func() {
		defer wg.Done()
		value := mu.load()
		fmt.Println("value go routine 3:", value)
	}()
	wg.Wait()
	fmt.Println("time elapsed:", time.Since(before))
}

type basicMutex struct {
	mu    sync.Mutex
	value int
}

func (m *basicMutex) store(value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}

func (m *basicMutex) load() int {
	m.mu.Lock()
	time.Sleep(time.Second)
	defer m.mu.Unlock()
	return m.value
}

type rwMutex struct {
	mu    sync.RWMutex
	value int
}

func (m *rwMutex) store(value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}

func (m *rwMutex) load() int {
	m.mu.RLock()
	time.Sleep(time.Second)
	defer m.mu.RUnlock()
	return m.value
}
