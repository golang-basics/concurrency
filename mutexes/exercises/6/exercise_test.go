package main

import (
	"testing"
	"time"
)

// GOFLAGS="-count=1" go test .
func TestExercise(t *testing.T) {
	now := time.Now()
	accounts := exercise()
	elapsed := time.Since(now)
	expectedAccounts := []account{
		{id: "a1", amount: 305, transactions: []string{"bank2: $100", "bank1: $200"}},
		{id: "a2", amount: 310, transactions: []string{"bank2: $100", "bank1: $200"}},
	}

	if len(accounts) != 2 {
		t.Fatalf("expected number of account to be 2, got %d", len(accounts))
	}
	for i, a := range accounts {
		if a.id != expectedAccounts[i].id || a.amount != expectedAccounts[i].amount || len(accounts[i].transactions) != 2 {
			t.Fatalf("expected account to be %v, got %v", expectedAccounts[i], a)
		}
	}
	if elapsed > 3*time.Second {
		t.Fatalf("exercise took too long: %v", elapsed)
	}
}
