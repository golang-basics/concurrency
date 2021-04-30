package pkg

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
