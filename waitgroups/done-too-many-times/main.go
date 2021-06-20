package main

import "sync"

func main() {
	var wg sync.WaitGroup
	// will result in panic
	wg.Done()
}
