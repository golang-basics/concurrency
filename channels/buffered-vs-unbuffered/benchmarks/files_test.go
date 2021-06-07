package benchmarks

import (
	"bufio"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"testing"
)

// to run the benchmark cd into "patterns/queuing/buffered-writing" and run:
// go test -bench=.
func BenchmarkUnbufferedWrite(b *testing.B) {
	b.ReportAllocs()
	write(b, newTmpFile())
}

// to run the benchmark cd into "patterns/queuing/buffered-writing" and run:
// go test -bench=.
func BenchmarkBufferedWrite(b *testing.B) {
	b.ReportAllocs()
	write(b, bufio.NewWriter(newTmpFile()))
}

func newTmpFile() *os.File {
	file, err := ioutil.TempFile("", "tmp")
	if err != nil {
		log.Fatalf("could not create temporary file: %v", err)
	}
	return file
}

func write(b *testing.B, w io.Writer) {
	done := make(chan struct{})
	defer close(done)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		_, err := w.Write([]byte(strconv.Itoa(i)))
		if err != nil {
			log.Fatalf("could not write to file: %v", err)
		}
	}
}
