package benchmarks

import (
	"runtime"
	"sync"
	"testing"
)

func lockedWorker(work, done chan struct{}) {
	runtime.LockOSThread()
	for {
		select {
		case <-done:
			runtime.UnlockOSThread()
			return
		case <-work:
		}
	}
}

func unlockedWorker(work, done chan struct{}) {
	for {
		select {
		case <-done:
			return
		case <-work:
		}
	}
}

func lockedWorkerSend(work, done, producerDone chan struct{}) {
	runtime.LockOSThread()
	for {
		select {
		case <-done:
			runtime.UnlockOSThread()
			return
		case <-work:
			producerDone <- struct{}{}
		}
	}
}

func producer(work chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	work <- struct{}{}
}

func producerReceive(work chan struct{}, producerDone chan struct{}, wg *sync.WaitGroup) {
	defer wg.Done()
	work <- struct{}{}
	<-producerDone
}

// go test -bench=.
func BenchmarkLockedWorker(b *testing.B) {
	work := make(chan struct{})
	done := make(chan struct{})
	defer close(done)
	go lockedWorker(work, done)

	var wg sync.WaitGroup
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Add(1)
			go producer(work, &wg)
		}
	})
	wg.Wait()
}

// go test -bench=.
func BenchmarkUnlockedWorker(b *testing.B) {
	work := make(chan struct{})
	done := make(chan struct{})
	defer close(done)
	go unlockedWorker(work, done)

	var wg sync.WaitGroup
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Add(1)
			go producer(work, &wg)
		}
	})
	wg.Wait()
}

// go test -bench=.
func BenchmarkLockedWorkerSendReceive(b *testing.B) {
	work := make(chan struct{})
	done := make(chan struct{})
	producerDone := make(chan struct{})
	defer close(done)
	go lockedWorkerSend(work, done, producerDone)

	var wg sync.WaitGroup
	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			wg.Add(1)
			go producerReceive(work, producerDone, &wg)
		}
	})
	wg.Wait()
}
