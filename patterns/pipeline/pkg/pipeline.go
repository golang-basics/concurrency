package pkg

type IntPipeline struct {
	out    chan int
	length int
}

func New(n int) *IntPipeline {
	p := IntPipeline{}
	out := make(chan int, n)
	for i := 1; i <= n; i++ {
		out <- i
	}
	p.length = n
	p.out = out
	return &p
}

func (p *IntPipeline) Increment() *IntPipeline {
	out := make(chan int, p.length)
	for i := 0; i < p.length; i++ {
		out <- <-p.out + 1
	}
	p.out = out
	return p
}

func (p *IntPipeline) Decrement() *IntPipeline {
	out := make(chan int, p.length)
	for i := 0; i < p.length; i++ {
		out <- <-p.out - 1
	}
	p.out = out
	return p
}

func (p *IntPipeline) Square() *IntPipeline {
	out := make(chan int, p.length)
	for i := 0; i < p.length; i++ {
		v := <-p.out
		out <- v * v
	}
	p.out = out
	return p
}

func (p *IntPipeline) Result() chan int {
	close(p.out)
	return p.out
}
