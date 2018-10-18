package main

import (
	"html/template"
	"net/http"
)

var tpl *template.Template

func init() {
	tpl = template.Must(template.ParseGlob("templates/*.gohtml"))//remember to make it and equal sign not a period
}

func main() {

	http.HandleFunc("/", def)
	http.HandleFunc("/about", abo)
	http.ListenAndServe(":8080", nil)
}

func def(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "default.gohtml", nil)
}
func abo(w http.ResponseWriter, r *http.Request) {
	tpl.ExecuteTemplate(w, "about.gohtml", nil)
}
