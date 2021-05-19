package main

import "fmt"

func main() {
	select {}
	fmt.Println("I never get printed")
}
