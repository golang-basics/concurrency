// The Circuit Breaker pattern is very useful for network (expensive calls).
// The idea of this pattern is very much similar like an electric circuit:
// when it is CLOSED (current flows through the circuit)
// when it is OPEN (no current flows through the circuit).
// The Circuit Breaker pattern introduces 3 states: CLOSED, OPEN and HALF-OPEN
// 1. The circuit is CLOSED for as long as the failure ratio is not high enough.
// 2. The circuit becomes OPEN for a given duration, when a certain error rate is met.
// 3. The circuit becomes HALF-OPEN after a given duration has elapsed.
// When the circuit if HALF-OPEN it will have a certain amount of retries
// If the retries fail, the circuit becomes OPEN again for the same duration.
// When the external API becomes healthy again, the circuit becomes CLOSED
// and all the requests flow seamlessly.
// For more info about the Circuit Breaker pattern feel free to check out:
// https://martinfowler.com/bliki/CircuitBreaker.html
package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/sony/gobreaker"
)

func main() {
	newHTTPServer()
	httpClient := newHTTPClient()

	var wg sync.WaitGroup
	wg.Add(40)
	for i := 0; i < 20; i++ {
		time.Sleep(time.Second)
		for j := 0; j < 2; j++ {
			go func() {
				defer wg.Done()
				_, _ = httpClient.Get("http://localhost:8080/bad")
			}()
		}
	}
	wg.Wait()
}

func newHTTPServer() {
	http.Handle("/bad", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// service is temporarily unavailable due to calls to external API
		// which for about 1 minute will be down
		fmt.Println("received an incoming request: bad")
		w.WriteHeader(http.StatusServiceUnavailable)
	}))
	http.Handle("/good", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// this route seems to be available, because it does not
		// depend on any external API
		fmt.Println("received an incoming request: good")
		w.WriteHeader(http.StatusOK)
	}))
	go func() {
		err := http.ListenAndServe(":8080", nil)
		if err != nil {
			log.Fatalf("http server error: %v", err)
		}
	}()
}

func newHTTPClient() httpClient {
	settings := gobreaker.Settings{
		MaxRequests: 2,
		Timeout:     5 * time.Second,
		ReadyToTrip: func(counts gobreaker.Counts) bool {
			return counts.ConsecutiveFailures > 5
		},
	}
	return httpClient{
		circuitBreaker: gobreaker.NewCircuitBreaker(settings),
	}
}

type httpClient struct {
	circuitBreaker *gobreaker.CircuitBreaker
}

func (client httpClient) Get(url string) ([]byte, error) {
	fmt.Println("current state:", client.circuitBreaker.State().String())
	body, err := client.circuitBreaker.Execute(func() (interface{}, error) {
		resp, err := http.Get(url)
		if err != nil {
			return nil, err
		}
		defer resp.Body.Close()

		if resp.StatusCode >= 500 {
			return nil, errors.New("we got a server error")
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}

		return body, nil
	})
	if err != nil {
		return nil, err
	}

	return body.([]byte), nil
}
