package main

import (
	"fmt"
)

func main() {
	m := map[string]string{"JD Salinger": "Catcher in The Rye", "Stephen King": "IT"}
	fmt.Printf("%T\n", m)
	fmt.Println(m["Stephen King"])
	fmt.Println(m)

	
}