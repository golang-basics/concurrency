package main

import (
	"net/http"
	"sync"
)
// To enable tracing on this program make sure to run the below command
// GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 go run main.go
func main() {
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			_, _ = http.Get("https://www.google.com")
			wg.Done()
		}()
	}
	wg.Wait()
}
