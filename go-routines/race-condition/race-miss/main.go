package main

import (
	"flag"
	"fmt"
	"sync"
)

// Try and run this program using the -race flag
// it will not display any race condition
// even if the code has a race condition.
// The race detector only works on running code.
// ---------------------------------------------
// The race condition will get detected only when
// the -inc and -dec flags are provided like so:
// go run -race -inc -dec
func main() {
	count := 0
	incrementFlag := flag.Bool("inc", false, "increment counter")
	decrementFlag := flag.Bool("dec", false, "decrement counter")
	flag.Parse()

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		if *incrementFlag {
			count++
		}
	}()
	go func() {
		defer wg.Done()
		if *decrementFlag {
			count--
		}
	}()
	wg.Wait()
	fmt.Println(count)
}
