package factorial

import (
	"testing"
)

func TestFactorial(t *testing.T) {
	testCases := []struct {
		name string
		num  uint
	}{
		{"0!", 0},
		{"1!", 1},
		{"10!", 10},
		{"20!", 20},
		{"55!", 55},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Factorial(tc.num)
			expect := factorial(tc.num)
			if actual != expect {
				t.Errorf("Factorial(%d) expect:%d, but actual:%d", tc.num, expect, actual)
			} else {
				t.Logf("Factorial(%d)=%d", tc.num, actual)
			}
		})
	}
}
