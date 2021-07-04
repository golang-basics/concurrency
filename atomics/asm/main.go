package main

import (
	"fmt"
)

func Abs(x float64) float64
func Sqrt(x float64) float64
func ASMWelcome()

func welcome() {
	fmt.Println("welcome")
}

func main() {
	ASMWelcome()
	fmt.Println(Abs(-12))
	fmt.Println(Sqrt(25))
}
