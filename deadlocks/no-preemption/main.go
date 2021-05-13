package main

func main() {
	done := make(chan struct{})
	go func() {
		for {
		}
		done <- struct{}{}
	}()
	go func() {
		<-done
		done <- struct{}{}
	}()
	<-done
}
