package main

import (
	"log"
	"net"
	"os"
	"sync"
	"testing"
)

func TestMain(m *testing.M) {
	srv := newTCPServer(":8080")
	defer srv.shutdown()
	os.Exit(m.Run())
}

// go test -bench=. -benchtime=3s
func BenchmarkTCPConnWithPool(b *testing.B) {
	b.ReportAllocs()
	pool := &sync.Pool{
		New: func() interface{} {
			conn, err := net.Dial("tcp", ":8080")
			if err != nil {
				log.Fatalf("could not dial: %v", err)
			}
			return conn
		},
	}
	pool.Put(pool.New())

	var wg sync.WaitGroup
	var mu sync.Mutex
	connects := 50
	for i := 0; i < b.N; i += connects {
		for j := 0; j < connects; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				conn := pool.Get().(*net.TCPConn)
				pool.Put(conn)

				mu.Lock()
				defer mu.Unlock()
				err := write(conn, "write")
				if err != nil {
					b.Errorf("error writing: %v", err)
				}
				_, err = read(conn)
				if err != nil {
					b.Errorf("error reading: %v", err)
				}
			}()
		}
		wg.Wait()
	}
}

// go test -bench=. -benchtime=3s
func BenchmarkTCPConnWithoutPool(b *testing.B) {
	b.ReportAllocs()
	var wg sync.WaitGroup
	var mu sync.Mutex
	connects := 50
	for i := 0; i < b.N; i += connects {
		for j := 0; j < connects; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				conn, err := net.Dial("tcp", ":8080")
				if err != nil {
					log.Fatalf("could not dial: %v", err)
				}

				mu.Lock()
				defer mu.Unlock()
				err = write(conn, "write")
				if err != nil {
					b.Errorf("error writing: %v", err)
				}
				_, err = read(conn)
				if err != nil {
					b.Errorf("error reading: %v", err)
				}

				err = write(conn, "close")
				if err != nil {
					b.Errorf("error closing: %v", err)
				}
			}()
		}
		wg.Wait()
	}
}
