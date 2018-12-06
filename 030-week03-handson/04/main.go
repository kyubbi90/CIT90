package main

import (
	"fmt"
)

type person struct {
	fname string
	lname string

}

func main() {
	p1 := person{
		"Eric",
		"Lopez",

	}
	fmt.Println(p1.fname)
}
