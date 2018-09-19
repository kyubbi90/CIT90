package main

import (
	"fmt"
)

type person struct {
	first    string
	last     string
	iceCream []string
}

func main() {
	p1 := person{
		first: "Eric",
		last:  "Lopez",
		iceCream: []string{
			"Chocolate",
			"vanilla",
			"cookies and cream",
		},
	}
	p2 := person{
		first: "James",
		last:  "Bond",
		iceCream: []string{
			"strawberry",
			"martnini",
		},
	}
	fmt.Println(p1.first)
	fmt.Println(p1.last)
	for i, v := range p1.iceCream{
		fmt.Println(i, v)
	}
	fmt.Println(p2.first)
	fmt.Println(p2.last)
	for i, v := range p2.iceCream{
		fmt.Println(i, v)
	}
}
