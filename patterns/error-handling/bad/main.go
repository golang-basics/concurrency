// You can find more about the better error handling in concurrent situations on:
// https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html

package main

import (
	"fmt"
	"net/http"
)

func main() {
	done := make(chan interface{})
	defer close(done)
	urls := []string{
		"https://www.google.com",
		"https://badhost1",
		"https://badhost2",
		"https://badhost3",
		"https://www.google.com",
		"https://www.google.com",
		"https://www.google.com",
	}

	for response := range checkStatus(done, urls...) {
		fmt.Println("response status:", response.Status)
	}
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan *http.Response {
	results := make(chan *http.Response)
	go func() {
		defer close(results)
		for _, url := range urls {
			response, err := http.Get(url)
			if err != nil {
				// let's hope someone will pay attention at the logged error
				// who's really responsible for handling the error?
				fmt.Println("error:", err)
				continue
			}
			select {
			case <-done:
				return
			case results <- response:
			}
		}
	}()
	return results
}
