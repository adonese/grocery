package main

import (
	"html/template"
	"net/http"
)

var temp, _ = template.ParseFiles("static/template.html")

func form(w http.ResponseWriter, r *http.Request) {
	data := map[string]string{
		"title": "Hello Solus",
	}
	temp.Execute(w, data)
}

func main() {
	http.HandleFunc("/", form)
	http.ListenAndServe(":8080", nil)
}
