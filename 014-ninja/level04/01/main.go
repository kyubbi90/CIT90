package main

import "fmt"

func main() {

	x := [5]int{10, 73, 9, 4, 26}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
