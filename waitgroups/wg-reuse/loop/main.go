package main

import (
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 100; i++ {
		// Imagine all go routines at once spawn
		// All at once add 3
		// couple of go routines call Wait and return fast
		// other go routines call Wait but return slower
		// other go routines call Wait while the slow go routines have not yet returned

		// the code below also has a data race
		go func() {
			wg.Add(3)
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			go func() {
				wg.Done()
			}()
			wg.Wait()
		}()
	}
}
