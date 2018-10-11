package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	//create string
	//assign to variable
	name := "Eric"
	str := `Html ` + name + ` html`
	fmt.Println(str)
	//create string
	//string print
	s := fmt.Sprint(`mas ` + name + ` menos`)
	fmt.Println(s)
	//create file
	//io.copy to the file
	nf, err := os.Create("newfile.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}
	io.Copy(nf, strings.NewReader(s))
	//create file
	//writestring to file
	nf2, err := os.Create("newfile2.txt")
	if err != nil {
		log.Fatal("whoops", err)
	}
	n, err := nf2.WriteString(str)
	if err != nil {
		log.Fatal("whoops2", err)
	}
	fmt.Println("bytes written:", n)

}
