package templates

const TakeTpl = `package {{.Pkg}}

func Take(done <-chan struct{}, in <-chan {{.Type}}, num int) <-chan {{.Type}} {
	out := make(chan {{.Type}})
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
}`
