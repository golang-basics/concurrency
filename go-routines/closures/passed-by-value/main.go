package main

import "fmt"

type person struct {
	name string
}

func main() {
	n := 10
	if n := ten(); n == 10 {
		fmt.Println("we got ten")
		n = 11
	}
	fmt.Println("n", n)
	tryToUpdateByValue()
	//tryToUpdateByIndex()
	//tryToUpdateByReference()
}

func tryToUpdateByValue() {
	people := []person{
		{name: "John"},
		{name: "Jane"},
		{name: "Amy"},
		{name: "Steve"},
	}

	fmt.Println("trying to update people names")
	for _, p := range people {
		fmt.Println(p.name)
		p.name = "Anonymous"
	}

	fmt.Println("checking if people names were updates")
	for _, p := range people {
		fmt.Println(p.name)
	}
}

func tryToUpdateByIndex() {
	people := []person{
		{name: "John"},
		{name: "Jane"},
		{name: "Amy"},
		{name: "Steve"},
	}

	fmt.Println("trying to update people names")
	for i, p := range people {
		fmt.Println(p.name)
		people[i].name = "Anonymous"
	}

	fmt.Println("checking if people names were updates")
	for _, p := range people {
		fmt.Println(p.name)
	}
}

func tryToUpdateByReference() {
	people := []*person{
		{name: "John"},
		{name: "Jane"},
		{name: "Amy"},
		{name: "Steve"},
	}

	fmt.Println("trying to update people names")
	for _, p := range people {
		fmt.Println(p.name)
		p.name = "Anonymous"
	}

	fmt.Println("checking if people names were updates")
	for _, p := range people {
		fmt.Println(p.name)
	}
}

func ten() int {
	return 10
}
