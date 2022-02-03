package main

import (
	"fmt"
	"sync"
	"time"
)

// EXERCISE:
// Find what's wrong in the exercise() function
// Make sure all the tests are passing
// DO NOT remove any Sleep() calls

// try running this with the -race flag
// go run -race exercise.go

// run the tests using:
// GOFLAGS="-count=1" go test .

func main() {
	exercise()
}

func exercise() {
	var mu sync.Mutex
	var n1, n2 int
	f1, f2 := file{mu: &mu, name: "f1"}, file{mu: &mu, name: "f2"}
	//write := func(f *file, name1, name2 string) {
	//	for i := 0; i < 5; i++ {
	//		if  f.write([]byte(name1)) || f.write([]byte(name2)) {
	//			return
	//		}
	//	}
	//}
	write := func(f *file, n1, n2 *int) {
		for i := 0; i < 5; i++ {
			if  f.write(n1) || f.write(n2) {
				fmt.Println(f.name, "successfully wrote to file")
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		write(&f1, &n1, &n2)
	}()
	go func() {
		defer wg.Done()
		write(&f2, &n1, &n2)
	}()

	wg.Wait()
	fmt.Println(n1)
	fmt.Println(n2)
}

type file struct {
	name string
	mu   *sync.Mutex
	//data []byte
}

//func (f *file) write(data []byte) bool {
func (f *file) write(n *int) bool {
	fmt.Println(f.name, "trying to write", *n)
	f.mu.Lock()
	//f.data = data
	*n += 1
	f.mu.Unlock()
	time.Sleep(500 * time.Millisecond)

	f.mu.Lock()
	//if bytes.Equal(f.data, data) {
	if *n == 1 {
		f.mu.Unlock()
		return true
	}
	f.mu.Unlock()


	time.Sleep(100 * time.Millisecond)

	f.mu.Lock()
	//f.data = nil
	*n -= 1
	f.mu.Unlock()
	time.Sleep(500 * time.Millisecond)
	return false
}
