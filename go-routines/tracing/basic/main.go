package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

// view trace after running the program by running the command
// go tool trace trace.out
func main() {
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}
	defer f.Close()
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()

	var wg sync.WaitGroup
	for i := 0; i < 30; i++ {
		wg.Add(1)
		go func() {
			t := 0
			for i := 0; i < 100; i++ {
				t += 2
			}
			fmt.Println("total:", t)
			wg.Done()
		}()
	}
	wg.Wait()
}
