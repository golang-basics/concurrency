package main

import (
	"fmt"
	"log"
	"net"
	"sync"
	"testing"
)

func init() {
	tcpServer(":8080").Wait()
}

func BenchmarkTCPConnWithPool(b *testing.B) {
	//var rLimit syscall.Rlimit
	//rLimit.Max = 99999
	//rLimit.Cur = 99999
	//err := syscall.Setrlimit(syscall.RLIMIT_NOFILE, &rLimit)
	//if err != nil {
	//	fmt.Println("could not set rlimit:", err)
	//}
	fmt.Println("N:", b.N)
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
	for i:=0;i<b.N;i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			//conn, err := net.Dial("tcp", ":8080")
			//if err != nil {
			//	log.Fatalf("could not dial: %v", err)
			//}
			conn := pool.Get().(*net.TCPConn)
			pool.Put(conn)

			mu.Lock()
			defer mu.Unlock()
			err := write(conn, "write")
			if err != nil {
				fmt.Printf("error writing: %v\n", err)
			}
			_, err = read(conn)
			if err != nil {
				fmt.Printf("error reading: %v\n", err)
			}
		}()
	}
	wg.Wait()
}
