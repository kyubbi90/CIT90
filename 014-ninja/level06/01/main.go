package main

import (
	"fmt"
)

func main() {
	n := foo()
	x, s := bar()

	fmt.Println(n, x, s)
}

func foo() int {
	return 27
}

func bar() (int, string) {
	return 21, "Test"
}