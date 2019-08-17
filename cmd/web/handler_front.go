package main

import (
	"html/template"
	"net/http"
)

// Data ...
type Data struct {
	Title string
}

func proccesTemplate(w http.ResponseWriter, page string) {
	w.Header().Set("Content-Type", "text/html")
	tmpl := template.Must(template.ParseFiles("template/layout/main.html", page))
	d := Data{Title: "Home Page"}
	tmpl.ExecuteTemplate(w, "main", d)
}

// Index home page
func Index(w http.ResponseWriter, r *http.Request) {
	proccesTemplate(w, "template/front/index.html")
}
