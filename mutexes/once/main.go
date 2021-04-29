package main

import (
	"fmt"
	"sync"
)

type Worker struct {
	once sync.Once
}

func (w *Worker) Run() {
	w.once.Do(func() {
		fmt.Println("this will run only once")
	})
}

func main() {
	w := Worker{}
	w.Run()
	w.Run()
	w.Run()
}
