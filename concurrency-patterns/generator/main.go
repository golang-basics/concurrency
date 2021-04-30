// The generator pattern is pretty simple and it resembles that
// there is a function which generates a channel of values
// and returns it at the end, while values are pushed by go routine(s)
// the channel must eventually close to avoid dead locks

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

func main() {
	for evenInt := range evenIntGen(5) {
		fmt.Println("even int:", evenInt)
	}
	for oddInt := range oddIntGen(5) {
		fmt.Println("odd int:", oddInt)
	}
	for word := range wordGen(5) {
		fmt.Println("word:", word)
	}
}

// generate random even integers
func evenIntGen(total int) <-chan int {
	out := make(chan int)
	go func() {
		n := randInt(10000)
		for i := n; i > n-total*2; i-- {
			if i%2 == 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

// generate random odd integers
func oddIntGen(total int) <-chan int {
	out := make(chan int)
	go func() {
		n := randInt(10000)
		for i := n; i > n-total*2; i-- {
			if i%2 != 0 {
				out <- i
			}
		}
		close(out)
	}()
	return out
}

// generate random words
func wordGen(total int) <-chan string {
	out := make(chan string)
	go func() {
		for i := 0; i < total; i++ {
			n := randInt(20)
			out <- strings.Repeat(string([]byte{65 + byte(n)}), n+1)
		}
		close(out)
	}()
	return out
}

func randInt(n int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	return r.Intn(n)
}
