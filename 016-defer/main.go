package main

import "fmt"

func main() {
defer foo()
bar()
}

func foo()  {
	fmt.Println("foooooo")
}
func bar()  {
	fmt.Println("barrrr")
	defer one()
	two()
}
func one()  {
	fmt.Println("one")
}
func two()  {
	fmt.Println("two")
}