package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string
	favfoods []string
}

func main() {
	p1 := person{
		"Eric",
		"Lopez",
		[]string{"fries", "burgers", "pizza"},

	}
	fmt.Println(p1.favfoods)
	for a, _ := range p1.favfoods {
		fmt.Printf(p1.favfoods[a])
	}
}
