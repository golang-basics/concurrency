package main

import (
	"log"
	"net"
	"sync/atomic"
	"time"
)

func main() {
	li, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("could not create listener: %v", err)
	}

	var connections int32
	for {
		conn, err := li.Accept()
		if err != nil {
			continue
		}
		connections++

		go func() {
			defer func() {
				_ = conn.Close()
				atomic.AddInt32(&connections, -1)
			}()
			if atomic.LoadInt32(&connections) > 3 {
				return
			}

			// simulate heavy work
			time.Sleep(time.Second)
			_, err := conn.Write([]byte("success"))
			if err != nil {
				log.Fatalf("could not write to connection: %v", err)
			}
		}()
	}
}
