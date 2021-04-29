package main

import (
	"fmt"
	"sync"
)

func main() {
	m := sync.Map{}
	// keys and values are interface{}
	m.Store("k1", "v1")
	m.Store("k2", "v2")
	m.Store("k3", "v3")

	value, ok := m.Load("k1")
	fmt.Println(value, ok)
	value, ok = m.Load("kn")
	fmt.Println(value, ok)

	m.Range(func(key, value interface{}) bool {
		fmt.Println("key:", key, "value:", value)
		// if we return false the range ends
		// return bool acts as a break
		return true
	})
	m.Delete("k1")
	fmt.Println(m.Load("k1"))
	// QUESTION: How does one get the length of sync.Map?
	// ANSWER: implement it yourself by using an atomic counter

	k2, loaded := m.LoadAndDelete("k2")
	fmt.Println("load and delete k2:", k2, loaded)
	k2, ok = m.Load("k2")
	fmt.Println("load k2", k2, ok)

	k2, loaded = m.LoadOrStore("k2", "v2")
	fmt.Println("load or store k2 1st:", k2, loaded)
	k2, loaded = m.LoadOrStore("k2", "v2")
	fmt.Println("load or store k2 2nd:", k2, loaded)
}
