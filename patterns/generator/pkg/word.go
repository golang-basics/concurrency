package pkg

import "strings"

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
