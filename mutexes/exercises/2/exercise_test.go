package main

import (
	"testing"
	"time"
)

// GOFLAGS="-count=1" go test .
func TestExercise(t *testing.T) {
	c := &catalog{data: map[string]string{
		"p1": "apples",
		"p2": "oranges",
		"p3": "grapes",
		"p4": "pineapple",
		"p5": "bananas",
	}}

	now := time.Now()
	exercise(c, "p1", "p2", "p3", "p4", "p5")
	elapsed := time.Since(now)
	numberOfProducts := len(c.data)

	if numberOfProducts != 1005 {
		t.Fatalf("wrong number of products in the catalog, got: %d", numberOfProducts)
	}
	if elapsed > 10*time.Millisecond {
		t.Fatalf("exercise took too long, elapsed: %v", elapsed)
	}
}
