package main

import (
	"html/template"
	"net/http"
)

func TemplateHandle(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, "Hello,world!")
}

func main() {
	http.HandleFunc("/template", TemplateHandle)
	http.ListenAndServe(":8080", nil)
}
