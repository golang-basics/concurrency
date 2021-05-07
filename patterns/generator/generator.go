package generator

import (
	"math/big"
	"math/rand"
	"strings"
	"time"
)

// EvenIntGen generates random even integers
func EvenIntGen(total int) <-chan int {
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

// OddIntGen generates random odd integers
func OddIntGen(total int) <-chan int {
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

// HexIntGen generates random hex numbers
func HexIntGen(total int) <-chan int {
	out := make(chan int)
	hex := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "A", "B", "C", "D", "E", "F"}
	go func() {
		for i := 0; i < total; i++ {
			n := randInt(len(hex) - 1)
			bigN := new(big.Int)
			i, wasSet := bigN.SetString(strings.Repeat(hex[n], n+1), 16)
			if wasSet {
				out <- int(i.Int64())
			}
		}
		close(out)
	}()
	return out
}

// WordGen generates random words
func WordGen(total int) <-chan string {
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
