package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Result string

type Searcher interface {
	Search(query string) Result
}

type YouTubeSearch struct {
}

func (YouTubeSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"youtube search result for: %q",
		query,
	))
}

type WebSearch struct {
}

func (WebSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"web search result for: %q",
		query,
	))
}

type ImageSearch struct {
}

func (ImageSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"image search result for: %q",
		query,
	))
}

type MapsSearch struct {
}

func (MapsSearch) Search(query string) Result {
	networkLatency()
	return Result(fmt.Sprintf(
		"maps search result for: %q",
		query,
	))
}

func main() {
	start := time.Now()
	results := google("golang")
	elapsed := time.Since(start)
	for _, res := range results {
		fmt.Println(res)
	}
	fmt.Println("elapsed", elapsed)
}

func google(query string) []Result {
	out := make(chan Result)
	replicas := []Searcher{
		WebSearch{},
		YouTubeSearch{},
		ImageSearch{},
		MapsSearch{},
	}
	// used for multiple search servers, not for normal index as you used here
	searchReplica := func(i int) {
		out <- replicas[i].Search(query)
	}
	for i := range replicas {
		go searchReplica(i)
	}

	results := make([]Result, 0)
	timeout := time.After(80 * time.Millisecond)
	for i := 0; i < len(replicas); i++ {
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

func first(query string, replicas ...Searcher) {
}

func networkLatency() {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	time.Sleep(time.Duration(r.Intn(100)) * time.Millisecond)
}
