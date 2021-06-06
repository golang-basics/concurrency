package main

import "fmt"

func main() {
	done := make(chan struct{})
	close(done)

	for num := range incDone(done, gen(done, 1, 2, 3, 4, 5, 6, 7, 8, 9)) {
		fmt.Println("num:", num)
	}
}

func gen(done chan struct{}, nums ...int) chan int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		for _, n := range nums {
			select {
			// even if you select done here
			// there's no guarantee that nothing was already written to out
			case <-done:
				return
			case out <- n:
			}
		}
	}()
	return out
}

func genDone(done chan struct{}, nums ...int) chan int {
	out := make(chan int, len(nums))
	go func() {
		defer close(out)
		var i int
		for {
			select {
			case <-done:
				return
			default:
				if i < len(nums) {
					i++
					out <- nums[i]
				}
			}
		}
	}()
	return out
}

func inc(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			select {
			// even if done is selected here, it does not mean
			// range might not have pulled another value from in
			// so there's no guarantee done is selected before a write to out
			case <-done:
				return
			case out <- i + 1:
			}
		}
	}()
	return out
}

func incDone(done chan struct{}, in <-chan int) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case v, ok := <-in:
				if ok {
					out <- v + 1
				}
			}
		}
	}()
	return out
}
