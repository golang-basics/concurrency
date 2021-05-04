// The cancellation pattern assumes each pipeline accept a done channel
// and listens to the done channel for early cancellation of producing new values.
// There are many ways to implement this approach, such as using the done channel idiom.
// Each pipeline has to listen to the cancellation channel and interrupt its work accordingly.
// Cancelling a pipeline explicitly will produce values partially or will stop right away
// As a best practice, it's better to call done using defer to avoid any hanging go routines
// and a predictable complete result.

package main

import (
	"fmt"
)

func main() {
	p := newIntPipeline(3)
	defer p.done()
	//for i := range p.inc().sq().done().dec().res() {
	for i := range p.inc().sq().dec().res() {
		fmt.Println(i)
	}
}

type intPipeline struct {
	outChan  chan int
	doneChan chan struct{}
	length   int
}

func newIntPipeline(n int) *intPipeline {
	p := &intPipeline{
		doneChan: make(chan struct{}),
		length:   n,
	}
	out := make(chan int, n)
	for i := 1; i <= n; i++ {
		out <- i
	}
	p.outChan = out
	return p
}

func (p *intPipeline) inc() *intPipeline {
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

func (p *intPipeline) dec() *intPipeline {
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

func (p *intPipeline) sq() *intPipeline {
	out := make(chan int, p.length)
	for i := 0; i < p.length; i++ {
		v := <-p.outChan
		select {
		case out <- v * v:
		case <-p.doneChan:
			return p
		}
	}
	p.outChan = out
	return p
}

func (p *intPipeline) done() *intPipeline {
	close(p.doneChan)
	return p
}

func (p *intPipeline) res() chan int {
	close(p.outChan)
	return p.outChan
}
