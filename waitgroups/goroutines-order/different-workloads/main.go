package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"time"
)

// To have a look at how things get scheduled run the following command:
// go build main.go
// GOMAXPROCS=1 GOGC=off GODEBUG=schedtrace=200,scheddetail=1 ./main
func main() {
	var wg sync.WaitGroup
	wg.Add(5)
	go task1(&wg)
	go task2(&wg)
	go task3(&wg)
	go task4(&wg)
	go task5(&wg)
	wg.Wait()
}

func task1(wg *sync.WaitGroup) {
	defer wg.Done()
	_, err := http.Get("http://localhost:8080")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("task 1: done")
}

func task2(wg *sync.WaitGroup) {
	defer wg.Done()
	var count int
	for i := 0; i < 1_000_000_000; i++ {
		count += i
	}
	fmt.Println("task 2: done")
}

func task3(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("task 3: done")
}

func task4(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(500 * time.Millisecond)
	fmt.Println("task 4: done")
}

func task5(wg *sync.WaitGroup) {
	defer wg.Done()
	time.Sleep(600 * time.Millisecond)
	fmt.Println("task 5: done")
}
