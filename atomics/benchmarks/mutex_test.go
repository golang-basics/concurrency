// cd into benchmarks
// go test -bench=.

package benchmarks

import (
	"fmt"
	"sync"
	"testing"
)

func BenchmarkMutexNumber(b *testing.B) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	var v int64
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				mu.Lock()
				v = int64(j)
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("v", v)
}

func BenchmarkMutexStruct(b *testing.B) {
	var wg sync.WaitGroup
	var mu sync.Mutex
	cfg := Config{}
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			for j := 0; j < b.N; j++ {
				mu.Lock()
				cfg = Config{
					a: []int{j + 1, j + 2, j + 3, j + 4, j + 5},
				}
				mu.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println("cfg", cfg)
}
