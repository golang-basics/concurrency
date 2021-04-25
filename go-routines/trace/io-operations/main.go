package main

import (
	"os"
	"sync"
)
// To enable tracing on this program make sure to run the below command
// GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 go run main.go
func main() {
	var wg sync.WaitGroup
	_, _ = os.Create("tmp.txt")
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			file, _ := os.Open("tmp.txt")
			_ = file.Close()
			wg.Done()
		}()
	}
	wg.Wait()
}
