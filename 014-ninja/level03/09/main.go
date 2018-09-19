package main

import (
	"fmt"
)

func main() {
	favSport := "football"
	switch favSport {
	case "football":
		fmt.Println("go to the field")
	case "baseball":
		fmt.Println("go to the baseball diamond")
	case "basketball":
		fmt.Println("go to the basketball courts")
	case "tennis":
		fmt.Println("go to the tennis courts")
	}
}
