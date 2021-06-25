package basic_testing

import (
	"testing"
)

func Sum(a, b int) int {
	return a + b
}

// All testing files must end with _test.go
// All test cases must be functions that begin with TestXxx
// The function signature is TestXxx(t *testing.T)
// To run all tests in current directory: go test .
// To run all tests in current directory, with detailed info: go test -v .

// Everything else in a test case is pretty much go code + some sugar functions provided by Go.
// You can also use a testing library such as testify with many more helpers & assertions
// https://github.com/stretchr/testify
func TestSum(t *testing.T) {
	res := Sum(1, 2)
	if res != 3 {
		t.Fatalf("expected %d to be %d", res, 3)
	}
}
