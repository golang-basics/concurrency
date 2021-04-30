package pkg

import (
	"math/big"
	"strings"
)

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
