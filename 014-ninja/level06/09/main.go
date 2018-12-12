package main

import "fmt"

func main() {

	y := func(xi []int) int {
		if len(xi) == 0 {
			return 0
		}
		if len(xi) == 1 {
			return xi[0]
		}
		return xi[0] + xi[len(xi)-1]
	}

	x := foo(y, []int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println(x)
}

func foo(y func(xi []int) int, ii []int) int {
	n := y(ii)
	n++
	return n
}
