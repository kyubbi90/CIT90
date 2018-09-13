package main

import "fmt"

var c int
var d = 43

type person struct{
	first string
	age int
	saying string
}
type secretagent struct{
	person
	ltk bool
}
func (sa secretagent) speak() {
	fmt.Println(sa.first, "says even more", sa.saying)
}
func (p person) speak(){
	fmt.Println(p.first, "says", p.saying)
}

func main() {
	a := "James"
	fmt.Println(a)
	fmt.Printf("%T\n", a)

	b := fmt.Sprint("Hello", a)
	fmt.Println(b)

	c = 42
	fmt.Println("c:", c)

	d = 43
	fmt.Println("d:", d)
	fmt.Printf("d type: %T\n", d)

	ee := []int{2, 3, 4, 7, 9}
	fmt.Println(ee)

	for i, v := range ee {
		fmt.Println(i, v)

	}

	f := map[string]int{"James":32, "Jenny":27,}
	fmt.Println(f)

	for k, v := range f{
		fmt.Println(k,v)
	}
	p1 := person{
		first:"James",
		age: 32,
		saying: "bleg",
	}
	p2 := person{
		first:"Jenny",
		age: 27,
		saying: "bleh",
	}
	xp := []person{p1,p2}
fmt.Println(xp)
for i2, v2 := range xp{
	fmt.Println(i2,v2)
}
for j := 0; j <10; j++ {
	fmt.Println(j)
}
sa1 := secretagent{
	person: person{
		first: "Jack",
		age: 29,
		saying: "blah",
	}
ltk: true
}
fmt.Println(sa1)
p1.speak()
p2.speak()
sa1.speak()
}
