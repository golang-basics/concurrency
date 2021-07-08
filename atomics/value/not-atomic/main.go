package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

type student struct {
	grades map[string]int
}

// to test this program, make sure to run it using the -race flag
// go run -race main.go
func main() {
	var wg sync.WaitGroup
	var val atomic.Value
	val.Store(student{grades: map[string]int{}})

	wg.Add(3)
	go func() {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["English"] = 10
		val.Store(student{grades: m})
		//s.grades["English"] = 10
	}()
	go func() {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["Math"] = 8
		val.Store(student{grades: m})
		//s.grades["Math"] = 8
	}()
	go func() {
		defer wg.Done()
		s := val.Load().(student)
		m := s.grades
		m["Physics"] = 7
		val.Store(student{grades: m})
		//s.grades["Physics"] = 7
	}()

	wg.Wait()
	s := val.Load().(student)
	fmt.Println(s)
}
