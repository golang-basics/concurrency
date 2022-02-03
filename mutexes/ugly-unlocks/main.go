package main

import (
	"fmt"
	"sync"
)

// try running this with the -race flag
// go run -race main.go
func main() {
	var count int
	var mu sync.Mutex
	var wg sync.WaitGroup

	wg.Add(1000)
	for i := 0; i < 1000; i++ {
		go func() {
			defer wg.Done()
			// ugly1(&mu, &count)
			// ugly2(&mu, &count)
			beautiful(&mu, &count)
		}()
	}

	wg.Wait()
	fmt.Println("count:", count)
}

func ugly1(mu *sync.Mutex, count *int) {
	mu.Lock()
	if *count > 5 {
		// It's so easy to forget to add this call to Unlock inside the if
		// causing a subtle deadlock, which can give you a headache
		// try commenting the below Unlock() :D
		// it will make the rest of waiting go routines to block on the first Lock()
		// thus infinitely waiting => deadlock.
		mu.Unlock()
		return
	}
	mu.Unlock()

	mu.Lock()
	*count++
	mu.Unlock()
}

func ugly2(mu *sync.Mutex, count *int) {
	// Now we don't have confusing Unlock() calls,
	// but we have extra variables, which can cause additional data problems if
	// the time it takes to manipulate the tmp data is too long.
	// We have to stay in the CRITICAL section for as long as the data is relevant.
	var tmp int
	mu.Lock()
	tmp = *count
	mu.Unlock()

	if tmp > 5 {
		return
	}

	mu.Lock()
	*count++
	mu.Unlock()
}

func beautiful(mu *sync.Mutex, count *int) {
	// because in our case we don't have a dedicated type,
	// plus we're reusing the mutex inside this function,
	// having a closure is the best, because we don't need
	// to pass redundant arguments like the mutex and count to it.
	// This way we keep the CRITICAL section safe and
	// avoid hassle and confusion by using a closure paired with defer.
	isGreater := func(n int) bool {
		mu.Lock()
		defer mu.Unlock()
		if *count > n {
			return true
		}
		return false
	}

	if isGreater(5) {
		return
	}

	mu.Lock()
	*count++
	mu.Unlock()
}
