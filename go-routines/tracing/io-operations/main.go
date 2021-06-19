package main

import (
	"os"
	"strconv"
	"sync"
	"time"
)

// To enable tracing on this program make sure to run the below commands
// go build main.go
// GOMAXPROCS=2 GODEBUG=schedtrace=1000,scheddetail=1 ./main
func main() {
	var wg sync.WaitGroup
	file, _ := os.Create("tmp.txt")
	defer file.Close()
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			_, err := file.WriteString(strconv.Itoa(n))
			if err != nil {
				panic(err)
			}
			time.Sleep(time.Second)
			wg.Done()
		}(i)
	}
	wg.Wait()

	// wait for Global Run Queue
	time.Sleep(3 * time.Second)
}
