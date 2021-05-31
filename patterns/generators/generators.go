// Having simple composable generators can be very powerful
// and can drastically simplify pipelines in Go
// For more info on useful fun generator check out this article:
// https://www.oreilly.com/library/view/concurrency-in-go/9781491941294/ch04.html

package generators

// Repeat generates an infinite number of repeated values till stopped via the done channel
func Repeat(done <-chan struct{}, values ...interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			for _, v := range values {
				select {
				case <-done:
					return
				case out <- v:
				}
			}
		}
	}()
	return out
}

// Take extracts a sub stream of values out of the value stream till it reaches a certain amount or stopped
func Take(done <-chan struct{}, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			case <-done:
				return
			case out <- <-in:
			}
		}
	}()
	return out
}

// RepeatFn works just like Repeat, but executes a function an infinite number of times till stopped
func RepeatFn(done <-chan struct{}, fn func() interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for {
			select {
			case <-done:
				return
			case out <- fn():
			}
		}
	}()
	return out
}

func ToInt(done <-chan struct{}, in <-chan interface{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
				return
			case out <- v.(int):
			}
		}
	}()
	return out
}
