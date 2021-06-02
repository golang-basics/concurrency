package main

import (
	"fmt"
	"strconv"
	"sync"
)

// try running the example like this: go run -race main.go
func main() {
	c := newChanCache()
	var wg sync.WaitGroup

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			c.set("k", "v"+strconv.Itoa(i))
		}(i)
	}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fmt.Println(c.get("k"))
		}()
	}
	wg.Wait()

	c.set("kn", "vn")
	fmt.Println(c.get("k"))
	fmt.Println(c.get("kn"))
}

type set struct {
	key   string
	value string
	done  chan struct{}
}

type get struct {
	key   string
	value chan string
}

type chanCache struct {
	cache   map[string]string
	setChan chan set
	getChan chan get
}

func newChanCache() chanCache {
	c := chanCache{
		cache:   map[string]string{},
		setChan: make(chan set),
		getChan: make(chan get),
	}
	go func() {
		for {
			select {
			case v := <-c.setChan:
				c.cache[v.key] = v.value
				v.done <- struct{}{}
			case v := <-c.getChan:
				v.value <- c.cache[v.key]
			}
		}
	}()
	return c
}

func (c chanCache) set(key, value string) {
	s := set{
		key:   key,
		value: value,
		done:  make(chan struct{}),
	}
	c.setChan <- s
	<-s.done
}

func (c chanCache) get(key string) string {
	g := get{
		key:   key,
		value: make(chan string),
	}
	c.getChan <- g
	return <-g.value
}
