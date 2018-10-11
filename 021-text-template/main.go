package main

import (
	"fmt"
	"os"
	"log"
	"text/template"
)

func main()  {
	tpl, err := template.ParseFiles("one.txt", "two.txt",)
	if err != nil {
		log.Fatal("whoops", err)
	}
		
	tpl.ExecuteTemplate(os.Stdout, "one.txt", nil)
	fmt.Println("\n -----")
	tpl.ExecuteTemplate(os.Stdout, "two.txt","Eric")
	fmt.Println("\n -----")
}