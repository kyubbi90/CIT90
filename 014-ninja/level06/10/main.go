package main

import "fmt"

func main() {
	y := foo()
	fmt.Println(y())
	fmt.Println(y())
	z := foo()
	fmt.Println(z())
	fmt.Println(z())
	fmt.Println(z())
	fmt.Println(z())
}

func foo() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}
