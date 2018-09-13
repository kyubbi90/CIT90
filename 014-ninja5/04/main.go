package main

import "fmt"


func main() {
	s := struct {
		first    string
		last     string
		iceCream []string
	}{
	
		first: "Eric",
		last:  "Lopez",
		iceCream: []string{
			"Chocolate",
			"vanilla",
			"cookies and cream",
		},
	}
	fmt.Println(s.first)
	fmt.Println(s.last)
	for k, v := range s.iceCream{
		fmt.Println(k,v)

	}
	
	
}