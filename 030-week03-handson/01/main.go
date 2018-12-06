package main

import (
	"fmt"
)

func main() {
	x := []int{2, 5, 50}
	for q, _ := range x  {
		fmt.Println(q)
	}
	for r, _ := range x  {
		fmt.Println(x[r])
	}
	for s, _ := range x {
		fmt.Println(s, "-", x[s])
	}
}
