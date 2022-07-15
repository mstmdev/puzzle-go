package flow

func Max(num1 int, num2 int) (result int)

// max the equivalent implementation of the Max function
func max(num1 int, num2 int) (result int) {
	if num1 >= num2 {
		result = num1
	} else {
		result = num2
	}
	return result
}

func Factorial(num uint) (result uint)

// factorial the equivalent implementation of the Factorial function
func factorial(num uint) (result uint) {
	result = 1
	if num == 0 {
		return result
	}
	for i := uint(1); i <= num; i++ {
		result *= i
	}
	return result
}
