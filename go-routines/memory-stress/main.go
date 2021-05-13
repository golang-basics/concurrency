package main

func main() {
	done := make(chan struct{})
	go func() {
		// memory allocations like this are pretty dangerous
		memory := make([]string, 0)
		for {
			memory = append(memory, "aaa")
		}
	}()

	<-done
}
