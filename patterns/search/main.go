// The search pattern resembles improvement on big/tail latency by
// using replicated servers and using the first value and discarding anything else
// Reduce tail latency using replicated search servers

package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	ip1 = "192.168.1.0"
	ip2 = "192.168.1.1"
)

func main() {
	start := time.Now()
	results := google("golang")
	elapsed := time.Since(start)
	for _, res := range results {
		fmt.Println(res)
	}
	fmt.Println("elapsed", elapsed)
}

type Result string

type Searcher interface {
	Search(query string) Result
}

type YouTubeSearch struct {
	ip string
}

func (s YouTubeSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"[%s] youtube search result for: %q",
		s.ip,
		query,
	))
}

type WebSearch struct {
	ip string
}

func (s WebSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"[%s] web search result for: %q",
		s.ip,
		query,
	))
}

type ImageSearch struct {
	ip string
}

func (s ImageSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"[%s] image search result for: %q",
		s.ip,
		query,
	))
}

type MapsSearch struct {
	ip string
}

func (s MapsSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"[%s] maps search result for: %q",
		s.ip,
		query,
	))
}

func google(query string) []Result {
	out := make(chan Result)

	go func() {
		out <- first(query, WebSearch{ip: ip1}, WebSearch{ip: ip2})
	}()
	go func() {
		out <- first(query, YouTubeSearch{ip: ip1}, YouTubeSearch{ip: ip2})
	}()
	go func() {
		out <- first(query, ImageSearch{ip: ip1}, ImageSearch{ip: ip2})
	}()
	go func() {
		out <- first(query, MapsSearch{ip: ip1}, MapsSearch{ip: ip2})
	}()

	results := make([]Result, 0)
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < 4; i++ {
		select {
		case res := <-out:
			results = append(results, res)
		case <-timeout:
			fmt.Println("timed out")
			return results
		}
	}

	return results
}

func first(query string, replicas ...Searcher) Result {
	out := make(chan Result)
	done := make(chan struct{})
	defer close(done)

	searchReplica := func(i int) {
		select {
		case out <- replicas[i].Search(query):
		case <-done:
		}
	}
	for i := range replicas {
		go searchReplica(i)
	}
	return <-out
}

func networkLatency() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	time.Sleep(time.Duration(r.Intn(100)) * time.Millisecond)
}
