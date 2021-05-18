package main

import (
	"fmt"
	"sync"
	"time"
)

type cart struct {
	cart []string
	cond *sync.Cond
	// add wait group for registration here
}

type orders struct {
	orders []string
	cond   *sync.Cond
}

type shipping struct {
	packages []string
	cond     *sync.Cond
}

func main() {
	c := cart{
		cart: make([]string, 0),
		cond: sync.NewCond(&sync.Mutex{}),
	}
	o := orders{
		orders: make([]string, 0),
		cond:   sync.NewCond(&sync.Mutex{}),
	}
	//s := shipping{
	//	packages: make([]string, 0),
	//	cond:     sync.NewCond(&sync.Mutex{}),
	//}

	var wg sync.WaitGroup
	wg.Add(1)
	go processShoppingCart(c, o, &wg)
	time.Sleep(time.Second)

	c.cart = append(c.cart, "jeans")
	c.cart = append(c.cart, "shirts")
	c.cart = append(c.cart, "sneakers")
	c.cond.Signal()

	wg.Wait()
}

func processShoppingCart(c cart, o orders, wg *sync.WaitGroup) {
	defer wg.Done()
	c.cond.L.Lock()
	c.cond.Wait()
	fmt.Println("processing shopping cart")
	for _, item := range c.cart {
		fmt.Println("processing:", item)
	}
	c.cond.L.Unlock()
}
