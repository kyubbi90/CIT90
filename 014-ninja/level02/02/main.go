package main

import "fmt"

func main()  {
	a := (14 == 14)
	b := (14 <= 20)
	c := (14 >= 20)
	d := (14 != 20)
	e := (14 < 20)
	f := (14 > 20)

	fmt.Println(a, b, c, d, e, f)
}