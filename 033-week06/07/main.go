package main

import (
	
	"log"
	"os"
	"text/template"
)
var tpl *template.Template
type person struct {
	first string
	last  string
}


func init() {
	tpl = template.Must(template.ParseFiles("index.gohtml"))
}



func main() {
	james := person{
		first: "James",
		last: "Bond",
	}
	miss := person{
		first: "Miss",
		last: "Moneypenny",
	}
	
	xp := []person{james, miss}
	


	err := tpl.Execute(os.Stdout, xp)
	if err != nil {
		log.Fatalln(err)
	}

}
