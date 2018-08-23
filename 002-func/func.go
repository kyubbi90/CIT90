package main

import (
	"fmt"
)

func main() {
	fmt.Println("hello world")
	secondStatement()
	finalStat()
}

//func receiver identifier(parameters) returns{code}

func secondStatement() {
	fmt.Println("here is my second statement")
}

func finalStat() {
	fmt.Println("about to exit")
}