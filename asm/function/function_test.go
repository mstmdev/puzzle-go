package function

import "testing"

func TestAdd(t *testing.T) {
	num1 := 11
	num2 := 22
	result := Add(num1, num2)
	t.Logf("%d+%d=%d", num1, num2, result)
}

func TestSub(t *testing.T) {
	num1 := 100
	num2 := 33
	result := Sub(num1, num2)
	t.Logf("%d-%d=%d", num1, num2, result)
}

func TestMul(t *testing.T) {
	num1 := 60
	num2 := 20
	result := Mul(num1, num2)
	t.Logf("%d*%d=%d", num1, num2, result)
}
