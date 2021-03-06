package main

import (
	"fmt"
)

type person struct {
	first  string
	last   string
	saying string
}

//func (receiver) identifier(parameters) (returns) {code}
func (p person) speak() {
	fmt.Println(p.first, "says", p.saying)
}

type human interface {
	speak()
}

func foo(h human) {
	h.speak()
}
func main() {
	p1 := person{
		first:  "James",
		last:   "Bond",
		saying: "Shaken, Not stirred",
	}
	p2 := person{
		first:  "Jenny",
		last:   "Moneypenny",
		saying: "Helllllo, James",
	}
	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println("------------")
	fmt.Println(p1, "says", p1.saying)
	fmt.Println(p2, "says", p2.saying)
	fmt.Println("------------")
	fmt.Println(p1.first, "says", p1.saying)
	fmt.Println(p2.first, "says", p2.saying)
	fmt.Println("------------")
	foo(p1)
	foo(p2)
}
