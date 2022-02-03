package main

import (
	"testing"
)

// GOFLAGS="-count=1" go test -race .
func TestExercise(t *testing.T) {
	clicks := exercise()

	n1 := clicks.elements["btn1"]
	n2 := clicks.elements["btn2"]
	total := clicks.total

	if n1 != 10 {
		t.Fatalf("expected number of clicks for btn1: %d, got: %d", 10, n1)
	}
	if n2 != 10 {
		t.Fatalf("expected number of clicks for btn2: %d, got: %d", 10, n2)
	}
	if total != 20 {
		t.Fatalf("expected total to be: %d, got: %d", 20, total)
	}
}
