package main

import (
	"fmt"
)

func main() {
	fmt.Println("begin")

	// pass in an ARGUMENT to the func when you CALL IT
	foo("James Bond")
	x:= "Miss Moneypenny"
	foo(x)
	n, _ := fmt.Println("1")
	fmt.Println("number of bytes written", n)
}

// foo takes a a VALUE of TYPE string
// foo has a parameter which is a VALUE of TYPE string
//funcs are defined with PARAMETER(S)
func foo(name string) {
	fmt.Println("hello", name)
}
