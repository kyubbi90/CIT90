package main

import "fmt"

type person struct {
	first   string
	age     int
	saying  string
	numbers []int
}

func (p person) speak() {
	fmt.Println(p.first, "says", p.saying)
}

func main() {

	p1 := person{
		first:   "James",
		age:     32,
		saying:  "bleh",
		numbers: []int{2, 3, 4, 7, 9},
	}

	for k, v := range p1.numbers {
		fmt.Println(k, v)
	}

	p1.speak()

	f := map[string]int{"James": 32, "Jenny": 27}

	for k, v := range f {
		fmt.Println(k, v)
	}

}
