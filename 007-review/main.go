package main

import (
	"fmt"
)
func main() {
	fmt.Println(`This is the entry point, beginning of the programming`)
	foo()
	a, b := bar("James Bond", 32)
	fmt.Println(a)
	fmt.Println(`In 10 years James Bond will be`, b, `years old` )
	fmt.Println(`Program "about" to exit`)
}
func foo() {
	fmt.Println("Foo is here")
}
func bar(x string, y int)  (string, int){
	return fmt.Sprint(x, " is here and he is ", y), y + 10
}
