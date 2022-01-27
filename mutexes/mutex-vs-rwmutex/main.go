package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	// mu := basicMutex{readSleepDuration: time.Second}
	mu := rwMutex{readSleepDuration: time.Second}
	mu.store(10)

	wg.Add(3)
	before := time.Now()
	go func() {
		defer wg.Done()
		value := mu.load()
		fmt.Println("value in go routine 1:", value)
	}()
	go func() {
		defer wg.Done()
		value := mu.load()
		fmt.Println("value in go routine 2:", value)
	}()
	go func() {
		defer wg.Done()
		value := mu.load()
		fmt.Println("value in go routine 3:", value)
	}()
	wg.Wait()
	fmt.Println("time elapsed:", time.Since(before))
}

type basicMutex struct {
	mu                sync.Mutex
	value             int
	readSleepDuration time.Duration
}

func (m *basicMutex) store(value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}

func (m *basicMutex) load() int {
	m.mu.Lock()
	time.Sleep(m.readSleepDuration)
	defer m.mu.Unlock()
	return m.value
}

type rwMutex struct {
	mu                sync.RWMutex
	value             int
	readSleepDuration time.Duration
}

func (m *rwMutex) store(value int) {
	m.mu.Lock()
	defer m.mu.Unlock()
	m.value = value
}

func (m *rwMutex) load() int {
	m.mu.RLock()
	time.Sleep(m.readSleepDuration)
	defer m.mu.RUnlock()
	return m.value
}
