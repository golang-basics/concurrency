// The server + worker pattern is based on the idea of the server and worker
// where the server is there to take care of each incoming connection
// and the worker(s) are responsible for easing and distributing the amount
// of work on multiple CPU cores, which might be too much for the server alone.
// Huge thanks to divan.dev. Check out the full resource here: https://divan.dev/posts/go_concurrency_visualize/

package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	addrCh := make(chan string)
	go server(l, addrCh)
	go pool(addrCh, 5)
	time.Sleep(10 * time.Second)
}

func server(l net.Listener, addrCh chan string) {
	for {
		conn, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(conn, addrCh)
	}
}

func handler(c net.Conn, addrCh chan string) {
	addr := c.RemoteAddr().String()
	addrCh <- addr
	time.Sleep(100 * time.Millisecond)
	_, _ = c.Write([]byte("ok"))
	_ = c.Close()
}

func pool(addrCh chan string, n int) {
	taskCh := make(chan int)
	for i := 0; i < n; i++ {
		go worker(taskCh)
	}
	for {
		addr := <-addrCh
		fmt.Println("running worker for addr:", addr)
		for i := 0; i < n; i++ {
			taskCh <- i
		}
	}
}

func worker(taskCh chan int) {
	for {
		fmt.Println("executing task:", <-taskCh)
	}
}
