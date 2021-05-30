package server

import (
	"fmt"
	"log"
	"net"
	"sync"
	"time"
)

func service() interface{} {
	time.Sleep(time.Second)
	return struct{}{}
}

func plainServer() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		srv, err := net.Listen("tcp", ":8080")
		if err != nil {
			log.Fatal("cannot listen:", err)
		}
		defer func() {
			_ = srv.Close()
		}()
		wg.Done()

		for {
			conn, err := srv.Accept()
			if err != nil {
				log.Printf("could not accept connection: %v", err)
				continue
			}

			service()
			_, _ = fmt.Fprintln(conn, "")
			_ = conn.Close()
		}
	}()
	return &wg
}

func serviceCacheWarmup() *sync.Pool {
	p := &sync.Pool{New: service}
	for i := 0; i < 10; i++ {
		p.Put(p.New())
	}
	return p
}

func poolServer() *sync.WaitGroup {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		srv, err := net.Listen("tcp", ":9090")
		if err != nil {
			log.Fatal("cannot listen:", err)
		}
		defer func() {
			_ = srv.Close()
		}()
		connPool := serviceCacheWarmup()
		wg.Done()

		for {
			conn, err := srv.Accept()
			if err != nil {
				log.Printf("could not accept connection: %v", err)
				continue
			}

			svcConn := connPool.Get()
			_, _ = fmt.Fprintln(conn, "")
			connPool.Put(svcConn)
			_ = conn.Close()
			// you will see a much slower performance
			// runtime.GC()
		}
	}()
	return &wg
}
