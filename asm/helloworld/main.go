package main

var hello = "hello world"

func main()

// printStr wrap the println function
// direct call runtime.printstring in main_amd64.s will cause compile error:
// relocation target runtime.printstring not defined for ABI0 (but is defined for ABIInternal)
func printStr(s string) {
	println(s)
}
