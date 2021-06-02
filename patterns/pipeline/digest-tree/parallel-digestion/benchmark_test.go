package main

import (
	"fmt"
	"testing"
)

// to run all the benchmarks cd into "digest-tree" directory and run
// go test -bench=. ./...
func BenchmarkMD5AllParallelDigestion(b *testing.B) {
	fmt.Println("parallel digestion")
}
