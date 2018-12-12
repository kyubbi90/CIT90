package main

import "fmt"

func main() {
	ii := []int{11, 12, 13, 14, 15, 16, 17, 18,19 }
	n := foo(ii...)
	fmt.Println(n)

	ii2 := []int{3, 4, 5, 6, 7, 8, 9, 10, 11}
	n2 := bar(ii2)
	fmt.Println(n2)
}

func foo(xi ...int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}

func bar(xi []int) int {
	total := 0
	for _, v := range xi {
		total += v
	}
	return total
}
