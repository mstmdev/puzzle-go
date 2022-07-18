package fibonacci

import (
	"testing"
)

func TestFibonacci(t *testing.T) {
	testCases := []struct {
		name string
		n    uint
	}{
		{"Fibonacci(1)", 1},
		{"Fibonacci(2)", 2},
		{"Fibonacci(3)", 3},
		{"Fibonacci(4)", 4},
		{"Fibonacci(5)", 5},
		{"Fibonacci(6)", 6},
		{"Fibonacci(7)", 7},
		{"Fibonacci(8)", 8},
		{"Fibonacci(9)", 9},
		{"Fibonacci(10)", 10},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			expect := fibonacci(tc.n)
			actual := Fibonacci(tc.n)
			if actual != expect {
				t.Errorf("Fibonacci(%d) expect:%d, but actual:%d", tc.n, expect, actual)
			} else {
				t.Logf("Fibonacci(%d)=%d", tc.n, actual)
			}
		})
	}
}
