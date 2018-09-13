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
			"martini",
		},
	}
	m := map[string]person{
		p1.last: p1,
		p2.last: p2,
	}

	for _, v := range m {
		fmt.Println(v.first)
		fmt.Println(v.last)
		for i, val := range v.iceCream {
			fmt.Println(i, val)
		}
		fmt.Println("---------")
	}
}
