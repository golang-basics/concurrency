package templates

const FanInTpl = `package {{.Pkg}}

import "sync"

func FanIn(done chan struct{}, inputs ...<-chan {{.Type}}) <-chan {{.Type}} {
	out := make(chan {{.Type}})
	var wg sync.WaitGroup
	wg.Add(len(inputs))

	for _, in := range inputs {
		go func(numbers <-chan {{.Type}}) {
			defer wg.Done()
			for n := range numbers {
				select {
				case <-done:
					return
				case out <- n:
				}
			}
		}(in)
	}
	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}`
