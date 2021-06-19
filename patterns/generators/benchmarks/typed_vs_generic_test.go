package benchmarks

import (
	"testing"

	"concurrency/patterns/generators"
)

var (
	toInt     = generators.ToInt
	repeat    = generators.Repeat
	intRepeat = generators.IntRepeat
	take      = generators.Take
	intTake   = generators.IntTake
)

func BenchmarkGenericGenerators(b *testing.B) {
	done := make(chan struct{})
	defer close(done)

	for i := 0; i < b.N; i++ {
		for range toInt(done, take(done, repeat(done, 1), 10)) {
		}
	}
}

func BenchmarkTypedGenerators(b *testing.B) {
	done := make(chan struct{})
	defer close(done)

	for i := 0; i < b.N; i++ {
		for range intTake(done, intRepeat(done, 1), 10) {
		}
	}
}
