package main

import (
	"html/template"
	"math/rand"
	"time"

	"net/http"
)

func DefineHandle(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	n := rand.Intn(10)
	var t *template.Template
	if n > 5 {
		t, _ = template.ParseFiles("layout.html", "red.html")
	} else {
		t, _ = template.ParseFiles("layout.html", "blue.html")
	}
	t.ExecuteTemplate(w, "layout", "")
}

func main() {
	http.HandleFunc("/define", DefineHandle)
	http.ListenAndServe(":8080", nil)
}
