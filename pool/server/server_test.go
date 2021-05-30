package server

import (
	"io/ioutil"
	"net"
	"testing"
)

func init() {
	// to avoid recreating servers that listen on the same address
	// inside each benchmark
	plainServer().Wait()
	poolServer().Wait()
}

// to run the benchmark run the following
// go test -benchtime=3s -bench=.
func BenchmarkPlainServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", ":8080")
		if err != nil {
			b.Fatalf("could not dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("could not read from connection: %v", err)
		}
		_ = conn.Close()
	}
}

// to run the benchmark run the following
// go test -benchtime=3s -bench=.
func BenchmarkPoolServer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		conn, err := net.Dial("tcp", ":9090")
		if err != nil {
			b.Fatalf("could not dial host: %v", err)
		}
		if _, err := ioutil.ReadAll(conn); err != nil {
			b.Fatalf("could not read from connection: %v", err)
		}
		_ = conn.Close()
	}
}
