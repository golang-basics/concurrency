package main

func main() {
	writeStream := make(chan<- int)
	readStream := make(<-chan int)

	// both statement above will result in compilation errors
	//<-writeStream
	//readStream<- 1

	// to avoid compilation errors
	writeStream = writeStream
	readStream = readStream
}
