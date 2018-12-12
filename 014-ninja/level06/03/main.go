package main

import "fmt"

func main() {
	defer foo()
	fmt.Println("Hello")
}

func foo() {
	defer func() {
		fmt.Println("foo DEFER ran")
	}()
	fmt.Println("foo ran")
}
