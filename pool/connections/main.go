package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"net"
	"sync"
)

const (
	addr               = ":8080"
	concurrentConnects = 30
)

// set max open files using: ulimit command
// ulimit -n 15000 && go run main.go
func main() {
	srv := newTCPServer(addr)
	defer srv.shutdown()
	var connectionsCreated int
	pool := &sync.Pool{
		New: func() interface{} {
			conn, err := net.Dial("tcp", addr)
			if err != nil {
				log.Fatalf("could not dial: %v", err)
			}
			connectionsCreated++
			return conn
		},
	}
	pool.Put(pool.New())

	var wg sync.WaitGroup
	var mu sync.Mutex
	for i := 0; i < 10000; i += concurrentConnects {
		for j := 0; j < concurrentConnects; j++ {
			wg.Add(1)
			go func() {
				defer wg.Done()
				conn := pool.Get().(*net.TCPConn)
				pool.Put(conn)

				mu.Lock()
				defer mu.Unlock()
				err := write(conn, "write")
				if err != nil {
					log.Printf("client: %v", err)
				}

				s, err := read(conn)
				if err != nil {
					log.Printf("client: %v", err)
				}
				fmt.Println("conn string:", s)
			}()
		}
		wg.Wait()
	}
	fmt.Println("connections created:", connectionsCreated)
}

type tcpServer struct {
	li net.Listener
}

func newTCPServer(addr string) *tcpServer {
	srv := tcpServer{}
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		li, err := net.Listen("tcp", addr)
		if err != nil {
			log.Fatalf("could not listen: %v", err)
		}
		srv.li = li
		wg.Done()

		for {
			conn, err := li.Accept()
			if err != nil {
				log.Printf("could not accept connection: %v", err)
				continue
			}
			go srv.serve(conn)
		}
	}()
	wg.Wait()
	return &srv
}

func (srv *tcpServer) serve(conn net.Conn) {
	scanner := bufio.NewScanner(conn)
	for scanner.Scan() {
		switch scanner.Text() {
		case "write":
			err := write(conn, "success")
			if err != nil {
				log.Printf("server write: %v", err)
			}
		case "close":
			err := write(conn, "closed")
			if err != nil {
				log.Printf("server close: %v", err)
			}
			_ = conn.Close()
			return
		}
	}
}

func (srv *tcpServer) shutdown() {
	_ = srv.li.Close()
}

func write(w io.Writer, s string) error {
	_, err := w.Write([]byte(s + "\n"))
	if err != nil {
		return fmt.Errorf("could not write to connection: %v", err)
	}
	return nil
}

func read(r io.Reader) (string, error) {
	reader := bufio.NewReader(r)
	bs, err := reader.ReadBytes('\n')
	if err != nil {
		return "", fmt.Errorf("could not read from connection: %v", err)
	}
	return string(bs[:len(bs)-1]), nil
}
