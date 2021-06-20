package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	wg.Wait()
	fmt.Println("Wait() executed immediately")
}
