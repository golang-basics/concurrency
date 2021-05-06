// The server pattern is based on the simple idea of infinite loop
// that listens for connections and runs each connection in a go routine.
// Huge thanks to divan.dev. Check out the full resource here: https://divan.dev/posts/go_concurrency_visualize/

package main

import "net"

func handler(c net.Conn) {
	_, _ = c.Write([]byte("ok"))
	_ = c.Close()
}

func main() {
	l, err := net.Listen("tcp", ":8080")
	if err != nil {
		panic(err)
	}
	for {
		c, err := l.Accept()
		if err != nil {
			continue
		}
		go handler(c)
	}
}
