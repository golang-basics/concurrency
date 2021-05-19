// cd into benchmarks
// go test -bench=.

package benchmarks

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

func BenchmarkAtomicNumber(b *testing.B) {
	var wg sync.WaitGroup
	var v int64
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				atomic.SwapInt64(&v, int64(j))
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("v", v)
}

func BenchmarkAtomicStruct(b *testing.B) {
	var wg sync.WaitGroup
	var v atomic.Value
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				v.Store(Config{
					a: []int{j + 1, j + 2, j + 3, j + 4, j + 5},
				})
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("cfg", v.Load())
}
