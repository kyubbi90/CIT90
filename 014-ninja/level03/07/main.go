package main

import (
	"fmt"
)

func main() {
	x := "Lopez"

	if x == "Lopez" {
		fmt.Println(x)
	} else if x == "Eric Lopez" {
		fmt.Println("EricEricEric", x)
	} else {
		fmt.Println("neither")
	}
}
