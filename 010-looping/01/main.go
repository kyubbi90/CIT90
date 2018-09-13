package main

import (
	"fmt"
)

func main() {
	/*  for{
		fmt.Print("hello")
	 } */

	for i := 0; i <= 10; i++ {
		fmt.Println(i)
	}
	//For statement needs 3 values or extra set of semicolon
	for i := 0; ; i++ {
		fmt.Println(i, i)
		if i == 11 {
			break
		}
	}
}
