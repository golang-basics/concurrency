package main

import (
	"testing"
	"time"
)

// GOFLAGS="-count=1" go test .
func TestExercise(t *testing.T) {
	now := time.Now()
	A := exercise()
	elapsed := time.Since(now)

	if A != 6000 {
		t.Fatalf("Expected A to be 6000, got: %d", A)
	}
	if elapsed > 10*time.Millisecond {
		t.Fatalf("exercise took too long: %v", elapsed)
	}
}
