package main

import (
	"fmt"
)

func main() {
	x := map[string]int{"Eric": 21, "Ariel": 24, "Juan": 51}

	fmt.Println(x)

	for k := range x {
		fmt.Println(k)
	}
	for v := range x {
		fmt.Println(x[v])
	}
	for k, v := range x {
		fmt.Println(k, ":", v)
	}
}
