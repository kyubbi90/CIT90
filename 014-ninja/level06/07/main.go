package main

import "fmt"

var x int
var y func()

func main() {

	z := func() {
		for i := 0; i < 3; i++ {
			fmt.Println(i)
		}
	}

	z()
	fmt.Printf("%T\n", z)

	fmt.Println(x)
	fmt.Printf("%T\n", x)

	y = z
	y()
	fmt.Printf("this is y %T\n", y)

	fmt.Println("done")
}
