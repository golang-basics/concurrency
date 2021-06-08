package main

import "fmt"

func main() {
	type foo int
	type bar int

	people := map[interface{}]string{
		0:               "Nobody",
		foo(0000):       "John",
		0 + 0i + bar(0): "Steve",
		int32(0):        "int32",
		int64(0):        "int64",
	}
	fmt.Println(people)
	fmt.Println(people[0000])
}
