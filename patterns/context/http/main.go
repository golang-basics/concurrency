package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"concurrency/patterns/context/mycontext"
)

const (
	server1Addr = ":8080"
	server2Addr = ":9090"
)

func main() {
	errChan := make(chan error)
	quitChan := make(chan os.Signal)
	signal.Notify(quitChan, syscall.SIGINT, os.Interrupt)

	go func() {
		fmt.Println("starting web server 1 on port", server1Addr)
		err := http.ListenAndServe(server1Addr, mux1())
		if err != nil {
			errChan <- err
		}
	}()
	go func() {
		fmt.Println("starting web server 2 on port", server2Addr)
		err := http.ListenAndServe(server2Addr, mux2())
		if err != nil {
			errChan <- err
		}
	}()

	select {
	case <-quitChan:
		fmt.Println("quiting")
	case err := <-errChan:
		fmt.Println("something bad happened:", err)
	}
}

func mux1() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", send("server1_value"))
	mux.HandleFunc("/receive", receive())
	return mux
}

func mux2() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/send", send("server2_value"))
	mux.HandleFunc("/receive", receive())
	return mux
}

func send(serverValue string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		ctx := mycontext.WithSomeValue(r.Context(), serverValue)
		ctx, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancel()
		req, err := http.NewRequestWithContext(ctx, http.MethodGet, "http://localhost:9090/receive", nil)
		if err != nil {
			log.Println("could not create request", err)
			return
		}
		res, err := http.DefaultClient.Do(mycontext.WithSomeValueRequest(req))
		if err != nil {
			log.Println("could not parse response", err)
			return
		}
		bs, err := ioutil.ReadAll(res.Body)
		if err != nil {
			log.Println("could not read response body", err)
			return
		}
		_, _ = w.Write(bs)
	}
}

func receive() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		//time.Sleep(time.Second)
		someValue := mycontext.SomeValueFromRequest(r)
		_, _ = w.Write([]byte(fmt.Sprintf("value: %s\n", someValue)))
	}
}
