package main

import (
	"fmt"
)

func Abs(float64) float64
func Sqrt(float64) float64
func ASMPrintResult(float64, float64)

func printResult(absRes, sqrtRes float64) {
	fmt.Println("abs result", absRes)
	fmt.Println("sqrt result", sqrtRes)
}

// to run this example run the following commands:
// go build -o exec
// ./exec
func main() {
	ASMPrintResult(Abs(-12), Sqrt(25))
}
