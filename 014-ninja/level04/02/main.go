package main

import "fmt"

func main() {
	x := []int{21, 54, 4, 27, 31, 98, 74, 23, 12, 89}
	for i, v := range x {
		fmt.Println(i, v)
	}
	fmt.Printf("%T\n", x)
}
