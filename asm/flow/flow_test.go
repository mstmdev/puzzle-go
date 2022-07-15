package flow

import (
	"testing"
)

func TestMax(t *testing.T) {
	testCases := []struct {
		name string
		num1 int
		num2 int
	}{
		{"less than", 10, 20},
		{"more than", 50, 30},
		{"equal", 60, 60},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := Max(tc.num1, tc.num2)
			expect := max(tc.num1, tc.num2)
			if actual != expect {
				t.Errorf("Max(%d,%d) expect:%d, but actual:%d", tc.num1, tc.num2, expect, actual)
			} else {
				t.Logf("Max(%d,%d)=%d", tc.num1, tc.num2, actual)
			}
		})
	}
}

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
