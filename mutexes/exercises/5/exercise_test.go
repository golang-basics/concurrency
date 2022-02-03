package main

import (
	"math"
	"testing"
	"time"
)

const delta = 5000

// GOFLAGS="-count=1" go test .
func TestExercise(t *testing.T) {
	p1, p2 := exercise(time.Second)

	min, max := math.Min(float64(p1), float64(p2)), math.Max(float64(p1), float64(p2))
	if max-min > delta {
		t.Fatalf("difference in values is bigger than delta: %d, got: %d", delta, int(max-min))
	}
}
