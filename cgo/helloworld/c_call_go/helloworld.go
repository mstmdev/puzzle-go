//go:build PUZZLE_CGO

package main

import "C"
import "fmt"

func main() {

}

//export hello
func hello() {
	fmt.Println("from golang: hello world")
}

//export say
func say(s string) {
	fmt.Printf("from golang: %s\n", s)
}
