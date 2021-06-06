package templates

const RepeatTpl = `package {{.Pkg}}

func Repeat(done <-chan struct{}, values ...{{.Type}}) <-chan {{.Type}} {
	out := make(chan {{.Type}})
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
}`
