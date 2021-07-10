package main

import (
	"fmt"
)

// https://github.com/golang/go/blob/master/src/math/sqrt.go#L92
// https://github.com/golang/go/blob/master/src/math/sqrt_asm.go#L12
// https://github.com/golang/go/blob/master/src/math/sqrt_amd64.s#L8
// Sqrt implementation is inside sqrt_amd64.s
func Sqrt(float64) float64

// https://github.com/golang/go/blob/master/src/math/floor.go#L13
// https://github.com/golang/go/blob/master/src/math/floor_asm.go#L12
// https://github.com/golang/go/blob/master/src/math/floor_amd64.s#L10
// Floor implementation is inside floor_amd64.s
func Floor(float64) float64

// Abs implementation is inside abs_amd64.s
func Abs(float64) float64

// ASMPrintResult implementation is inside print_amd64.s
func ASMPrintResult(float64, float64, float64)

func printResult(absRes, sqrtRes, floorRes float64) {
	fmt.Println("abs result", absRes)
	fmt.Println("sqrt result", sqrtRes)
	fmt.Println("floor result", floorRes)
}

// to run this example run the following commands:
// go build -o exec
// ./exec
func main() {
	ASMPrintResult(Abs(-12), Sqrt(25), Floor(2.56))
}
