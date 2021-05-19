package main

import (
	"fmt"
	"sync"
	"time"
)

type cart struct {
	cart []string
	cond *sync.Cond
	wg   *sync.WaitGroup
}

type orders struct {
	orders []string
	cond   *sync.Cond
	wg     *sync.WaitGroup
}

type shipping struct {
	packages []string
	cond     *sync.Cond
	wg       *sync.WaitGroup
}

func main() {
	shopping := &sync.WaitGroup{}
	c := &cart{
		cart: make([]string, 0),
		cond: sync.NewCond(&sync.Mutex{}),
		wg:   shopping,
	}
	o := &orders{
		orders: make([]string, 0),
		cond:   sync.NewCond(&sync.Mutex{}),
		wg:     shopping,
	}
	s := &shipping{
		packages: make([]string, 0),
		cond:     sync.NewCond(&sync.Mutex{}),
		wg:       shopping,
	}

	var steps sync.WaitGroup
	steps.Add(3)
	go processShoppingCart(c, o, &steps)
	go processOrders(o, s, &steps)
	go processShipping(s, &steps)
	steps.Wait()

	c.cart = append(c.cart, "jeans")
	c.cart = append(c.cart, "shirts")
	c.cart = append(c.cart, "sneakers")
	c.cond.Signal()

	shopping.Wait()

	whLookup := &sync.WaitGroup{}
	whCond := sync.NewCond(&sync.Mutex{})
	wh1 := warehouse{id: "1", wg: whLookup, cond: whCond}
	wh2 := warehouse{id: "2", wg: whLookup, cond: whCond}
	wh3 := warehouse{id: "3", wg: whLookup, cond: whCond}

	steps.Add(3)
	go wh1.process(s.packages, &steps)
	go wh2.process(s.packages, &steps)
	go wh3.process(s.packages, &steps)
	steps.Wait()

	fmt.Println("PACKAGE LOOKUP")
	whCond.Broadcast()
	whLookup.Wait()
}

func processShoppingCart(c *cart, o *orders, wg *sync.WaitGroup) {
	c.wg.Add(1)
	defer c.wg.Done()

	wg.Done()
	c.cond.L.Lock()
	c.cond.Wait()
	fmt.Println("PROCESSING SHOPPING CART")
	for _, item := range c.cart {
		fmt.Println("processing cart:", item)
		// simulate work and network
		time.Sleep(200 * time.Millisecond)
		o.orders = append(o.orders, item)
	}
	fmt.Println("---------")
	c.cond.L.Unlock()
	o.cond.Signal()
}

func processOrders(o *orders, s *shipping, wg *sync.WaitGroup) {
	o.wg.Add(1)
	defer o.wg.Done()

	wg.Done()
	o.cond.L.Lock()
	o.cond.Wait()
	fmt.Println("PROCESSING ORDERS")
	for _, order := range o.orders {
		fmt.Println("processing order:", order)
		// simulate work and network
		time.Sleep(200 * time.Millisecond)
		s.packages = append(s.packages, order)
	}
	fmt.Println("---------")
	o.cond.L.Unlock()
	s.cond.Signal()
}

func processShipping(s *shipping, wg *sync.WaitGroup) {
	s.wg.Add(1)
	defer s.wg.Done()

	wg.Done()
	s.cond.L.Lock()
	s.cond.Wait()
	fmt.Println("PROCESSING SHIPPING")
	for _, pkg := range s.packages {
		// simulate work and network
		time.Sleep(200 * time.Millisecond)
		fmt.Println("processing package:", pkg)
	}
	fmt.Println("---------")
	s.cond.L.Unlock()
}

type warehouse struct {
	id   string
	cond *sync.Cond
	wg   *sync.WaitGroup
}

func (wh warehouse) process(packages []string, wg *sync.WaitGroup) {
	wh.wg.Add(1)
	defer wh.wg.Done()

	wg.Done()
	wh.cond.L.Lock()
	wh.cond.Wait()
	fmt.Println("warehouse", wh.id)
	for _, pkg := range packages {
		// simulate work and network
		time.Sleep(200 * time.Millisecond)
		fmt.Println("looking for package:", pkg)
	}
	fmt.Println("---------")
	wh.cond.L.Unlock()
}
