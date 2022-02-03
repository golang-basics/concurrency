package main

import (
	"testing"
)

// GOFLAGS="-count=1" go test .
func TestExercise(t *testing.T) {
	count := exercise()
	if count != 999000 {
		t.Fatalf("got a different count: %d", count)
	}
}
