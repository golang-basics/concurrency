package main

import "sync"

func main() {
	var mu sync.Mutex
	mu.Lock()
	defer mu.Unlock()

	mu.Lock()
	defer mu.Unlock()
}
