package main

import (
	"html/template"
	"math/rand"
	"net/http"
	"time"
)

func ConditionHandle(w http.ResponseWriter, r *http.Request) {
	rand.Seed(time.Now().Unix())
	ra := rand.Intn(10)
	t, _ := template.ParseFiles("tmpl.html")
	t.Execute(w, ra > 5)
}

func RangeHandle(w http.ResponseWriter, r *http.Request) {
	slice := []string{"MON", "TUES", "WED"}
	t, _ := template.ParseFiles("range.html")
	t.Execute(w, slice)

}

func main() {
	http.HandleFunc("/condition", ConditionHandle)
	http.HandleFunc("/range", RangeHandle)
	http.ListenAndServe(":8080", nil)
}
