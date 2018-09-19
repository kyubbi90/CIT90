package main

import (
	"fmt"
)

func main() {
	bd := 1997
	for {
		if bd > 2018 {
			break
		}
		fmt.Println(bd)
		bd++
	}
}
