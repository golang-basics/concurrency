package main

func main() {
	c := make(chan int)
	close(c)
	// this gets ignored
	<-c
	// this results in an error
	c <- 1
}
