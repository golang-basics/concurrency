package ctxswitch

import (
	"sync"
	"testing"
)

func BenchmarkContextSwitch(b *testing.B) {
	var wg sync.WaitGroup
	begin := make(chan struct{})
	c := make(chan struct{})
	sender := func() {
		defer wg.Done()
		for i := 0; i < b.N; i++ {
			c <- struct{}{}
		}
	}
	receiver := func() {
		defer wg.Done()
		<-begin
		for i := 0; i < b.N; i++ {
			<-c
		}
	}

	wg.Add(2)
	go sender()
	go receiver()

	b.StartTimer()
	close(begin)

	wg.Wait()
}
