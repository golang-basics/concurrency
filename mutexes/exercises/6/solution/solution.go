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

// try running this with the -race flag
// go run -race exercise.go

// run the tests using:
// GOFLAGS="-count=1" go test .

// SOLUTION
// The problem we are trying to solve here is a Livelock
// and that's a really hard to debug problem. The problem
// with this code lies in the transfer function, which causes
// each go routine to constantly reset the state, resulting in
// exhausting all available retries. The fix is pretty simple:
// All we need to do is correct the if statement and use a
// temporary variable like below.
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
	var accAmount float64
	b.mu.Lock()
	acc.amount += amount
	accAmount = acc.amount
	b.mu.Unlock()
	time.Sleep(500 * time.Millisecond)

	// check if the money were transferred
	b.mu.Lock()
	// if current balance equals to previous account balance
	if acc.amount == accAmount {
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
