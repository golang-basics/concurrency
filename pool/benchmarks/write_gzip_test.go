package benchmarks

import (
	"bytes"
	"compress/gzip"
	"io"
	"io/ioutil"
	"sync"
	"testing"
)

const lorem = "Lorem ipsum dolor sit amet, consectetur adipiscing elit. Pellentesque molestie."

func BenchmarkWriteGzipWithPool(b *testing.B) {
	pool := sync.Pool{
		New: func() interface{} {
			return gzip.NewWriter(ioutil.Discard)
		},
	}
	b.ResetTimer()
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := bytes.NewReader([]byte(lorem))
		writer := pool.Get().(*gzip.Writer)
		_ = writer.Flush()
		_, _ = io.Copy(writer, data)
		pool.Put(writer)
	}
}

func BenchmarkWriteGzipWithoutPool(b *testing.B) {
	b.ReportAllocs()
	for n := 0; n < b.N; n++ {
		data := bytes.NewReader([]byte(lorem))
		writer := gzip.NewWriter(ioutil.Discard)
		_, _ = io.Copy(writer, data)
	}
}
