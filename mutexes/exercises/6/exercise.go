package main

import (
	"fmt"
	"sync"
	"time"
)

// EXERCISE:
// Find what's wrong in the exercise() function
// Make sure all the tests are passing
// DO NOT remove any Sleep() calls
// DO NOT remove any steps, i.e: transfer, check, revert

// try running this with the -race flag
// go run -race exercise.go

// run the tests using:
// GOFLAGS="-count=1" go test .

func main() {
	now := time.Now()
	accounts := exercise()
	for _, a := range accounts {
		fmt.Printf("'%s' amount: $%v\n", a.id, a.amount)
		fmt.Printf("'%s' transactions: %v\n", a.id, a.transactions)
	}
	fmt.Println("elapsed:", time.Since(now))
}

func exercise() []*account {
	var mu sync.Mutex
	accounts := []*account{
		{id: "a1", amount: 5},
		{id: "a2", amount: 10},
	}
	b1 := bank{mu: &mu, name: "bank1"}
	b2 := bank{mu: &mu, name: "bank2"}
	write := func(b *bank, amount float64, accounts ...*account) {
		for i := 0; i < 5; i++ {
			ok := true
			for _, a := range accounts {
				if !b.transfer(a, amount) {
					ok = false
					break
				}
			}
			if ok {
				fmt.Printf("'%s' successfully transferred: $%v\n", b.name, amount)
				return
			}
		}
	}

	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		// transfer $200 from bank1 to all accounts
		write(&b1, 200, accounts...)
	}()
	go func() {
		defer wg.Done()
		// transfer $100 from bank2 to all accounts
		write(&b2, 100, accounts...)
	}()

	wg.Wait()
	return accounts
}

type account struct {
	id           string
	amount       float64
	transactions []string
}

type bank struct {
	name string
	mu   *sync.Mutex
}

func (b *bank) transfer(acc *account, amount float64) bool {
	// transfer the money
	fmt.Printf("'%s' trying to transfer: $%v to '%s'\n", b.name, amount, acc.id)
	b.mu.Lock()
	acc.amount += amount
	b.mu.Unlock()
	time.Sleep(500 * time.Millisecond)

	// check if the money were transferred
	b.mu.Lock()
	// if current balance equals to previous account balance
	if acc.amount == acc.amount-amount {
		b.mu.Unlock()
		tx := fmt.Sprintf("%s: $%v", b.name, amount)
		acc.transactions = append(acc.transactions, tx)
		return true
	}
	b.mu.Unlock()

	time.Sleep(100 * time.Millisecond)

	// revert the transfer
	b.mu.Lock()
	acc.amount -= amount
	b.mu.Unlock()
	time.Sleep(500 * time.Millisecond)
	return false
}
