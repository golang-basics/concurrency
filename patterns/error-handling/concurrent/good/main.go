// The best practice when it comes to error handling in a concurrent world
// is to tie them to responses, or wrap them inside a struct where both: the result and the error
// is present, this way each go routine would only care about sending a struct
// and let the caller decide better how to handle the error in case there is one.
// You can find more about the better error handling in concurrent situations on:
// https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html

package main

import (
	"fmt"
	"net/http"
)

type Result struct {
	Error    error
	Response *http.Response
}

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
	for result := range checkStatus(done, urls...) {
		// how many errors are too many errors?
		if result.Error != nil {
			fmt.Println("error:", result.Error)
			continue
		}
		fmt.Println("response status:", result.Response.Status)
	}
}

func checkStatus(done <-chan interface{}, urls ...string) <-chan Result {
	results := make(chan Result)
	go func() {
		defer close(results)
		for _, url := range urls {
			var result Result
			// instead of separating the error from the response
			// we couple the 2 together thus having a corresponding error for every result
			// making things easier to handle on the higher level
			response, err := http.Get(url)
			result = Result{
				Error:    err,
				Response: response,
			}
			select {
			case <-done:
				return
			case results <- result:
			}
		}
	}()
	return results
}
