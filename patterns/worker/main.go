// The worker pattern is very similar to the fan-out pattern where
// multiple go routines (workers) can read from the same channel, thus
// distributing the amount of work between the CPU cores.
// Huge thanks to divan.dev. Check out the full resource here: https://divan.dev/posts/go_concurrency_visualize/

package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

const (
	numberOfWorkers    = 5
	numberOfSubWorkers = 3
	numberOfTasks      = 20
	numberOfSubTasks   = 10
)

func main() {
	var wg sync.WaitGroup
	wg.Add(numberOfWorkers)
	tasks := make(chan int)

	for i := 0; i < numberOfWorkers; i++ {
		go worker(tasks, &wg)
	}

	for i := 0; i < numberOfTasks; i++ {
		tasks <- i
	}

	close(tasks)
	wg.Wait()
}

func worker(tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		task, ok := <-tasks
		if !ok {
			return
		}

		fmt.Println("executing task:", task)
		subTasks := make(chan int)
		for i := 0; i < numberOfSubWorkers; i++ {
			go subWorker(subTasks)
		}
		for i := 0; i < numberOfSubTasks; i++ {
			subTask := task * i
			subTasks <- subTask
		}
		close(subTasks)
	}
}

func subWorker(subtasks chan int) {
	for {
		subTask, ok := <-subtasks
		if !ok {
			return
		}
		rand.Seed(time.Now().UnixNano())
		time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
		fmt.Println("executing subtask:", subTask)
	}
}
