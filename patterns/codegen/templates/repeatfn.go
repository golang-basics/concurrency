package templates

const RepeatFnTpl = `package {{.Pkg}}

func RepeatFn(done <-chan struct{}, fn func() {{.Type}}) <-chan {{.Type}} {
	out := make(chan {{.Type}})
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
}`
