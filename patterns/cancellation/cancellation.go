// The cancellation pattern assumes each pipeline accept a done channel
// and listens to the done channel for early cancellation of producing new values.
// There are many ways to implement this approach, such as using the done channel idiom.
// Each pipeline has to listen to the cancellation channel and interrupt its work accordingly.
// Cancelling a pipeline explicitly will produce values partially or will stop right away
// As a best practice, it's better to call done using defer to avoid any hanging go routines
// and a predictable complete result.

package cancellation

type IntPipeline struct {
	outChan  chan int
	doneChan chan struct{}
	length   int
}

func NewIntPipeline(vs ...int) *IntPipeline {
	p := &IntPipeline{
		doneChan: make(chan struct{}),
		length:   len(vs),
	}
	out := make(chan int, len(vs))
	for _, n := range vs {
		out <- n
	}
	p.outChan = out
	return p
}

func (p *IntPipeline) Inc() *IntPipeline {
	out := make(chan int, p.length)
	for i := 0; i < p.length; i++ {
		select {
		case out <- <-p.outChan + 1:
		case <-p.doneChan:
			return p
		}
	}
	p.outChan = out
	return p
}

func (p *IntPipeline) Dec() *IntPipeline {
	out := make(chan int, p.length)
	for i := 0; i < p.length; i++ {
		select {
		case out <- <-p.outChan - 1:
		case <-p.doneChan:
			return p
		}
	}
	p.outChan = out
	return p
}

func (p *IntPipeline) Sq() *IntPipeline {
	out := make(chan int, p.length)
	for i := 0; i < p.length; i++ {
		select {
		case v := <-p.outChan:
			out <- v * v
		case <-p.doneChan:
			return p
		}
	}
	p.outChan = out
	return p
}

func (p *IntPipeline) Done() *IntPipeline {
	close(p.doneChan)
	return p
}

func (p *IntPipeline) Res() chan int {
	close(p.outChan)
	return p.outChan
}

func Gen(done chan struct{}, vs ...int) chan int {
	out := make(chan int, len(vs))
	defer close(out)
	for _, n := range vs {
		select {
		case <-done:
		case out <- n:
		}
	}
	return out
}

func Inc(done chan struct{}, in <-chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
			case out <- v + 1:
			}
		}
	}()
	return out
}

func Dec(done chan struct{}, in <-chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
			case out <- v - 1:
			}
		}
	}()
	return out
}

func Sq(done chan struct{}, in <-chan int) chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for v := range in {
			select {
			case <-done:
			case out <- v * v:
			}
		}
	}()
	return out
}
