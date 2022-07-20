package main

// #include "helloworld.h"
import "C"

func main() {
	C.hello()
	C.say(C.CString("bye bye\n"))
}
