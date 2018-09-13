package main

import "fmt"

func main()  {
	foo()

	func (x int){
		fmt.Println(x)
	}(2)
}
func foo()  {
	fmt.Println("fooooooo")
}