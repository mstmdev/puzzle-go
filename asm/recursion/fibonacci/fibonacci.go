package fibonacci

// Fibonacci a recursion implementation for the fibonacci sequence
func Fibonacci(n uint) (result uint)

// fibonacci the equivalent implementation of the Fibonacci function
func fibonacci(n uint) (result uint) {
	if n == 0 {
		return 0
	}
	if n <= 2 {
		return 1
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
