package main

import (
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkChannel(b *testing.B) {
	b.ReportAllocs()
	var value int32
	setChan := make(chan set)

	go func() {
		for {
			select {
			case op := <-setChan:
				value = op.value
				op.done <- struct{}{}
			}
		}
	}()
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			op := set{
				value: i,
				done:  make(chan struct{}),
			}
			setChan <- op
			<-op.done
		}(int32(i))
	}

	wg.Wait()
	// to avoid compile errors
	value = value
}

func BenchmarkAtomics(b *testing.B) {
	b.ReportAllocs()
	var value int32
	var wg sync.WaitGroup
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int32) {
			defer wg.Done()
			atomic.StoreInt32(&value, i)
		}(int32(i))
	}
	wg.Wait()
}
