package recursion

// Factorial a recursion implementation for factorial
func Factorial(num uint) (result uint)

// factorial the equivalent implementation of the Factorial function
func factorial(num uint) (result uint) {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)
}
