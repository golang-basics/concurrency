// To run all the benchmarks inside the current directory:
// go test -bench=.
// To run all the benchmarks inside the current directory for some time:
// go test -bench=. -benchtime=3s
package basic_testing

import (
	"testing"
)

func EfficientSum(a, b int) int {
	return a + b
}

func InefficientSum(a, b int) int {
	res := make(chan int, 1)
	res <- a + b
	return <-res
}

// Every benchmark must be stored in a file with the extension _test.go
// Every benchmark is a function which begins with BenchmarkXxx
// Every benchmark function has the signature BenchmarkXxx(b *testing.B)
func BenchmarkEfficientSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		EfficientSum(i, i+1)
	}
}

func BenchmarkInefficientSum(b *testing.B) {
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		InefficientSum(i, i+1)
	}
}
