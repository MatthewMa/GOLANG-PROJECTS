package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Post struct {
	User    string
	Threads []string
}

func WelcomeHandle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome!")
}

func RequestHandle(w http.ResponseWriter, r *http.Request) {
	//r.ParseForm()
	//r.ParseMultipartForm(1024)
	//fileHeader := r.MultipartForm.File["uploaded"][0]
	//file, err := fileHeader.Open()
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func BodyHandle(w http.ResponseWriter, r *http.Request) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	fmt.Fprintln(w, string(body))
}

func WriteHandle(w http.ResponseWriter, r *http.Request) {
	str := `<html>
	<head><title>Go Web Programming</title></head>
	<body><h1>Hello World</h1></body>
	</html>`
	w.Write([]byte(str))
}

func WriteHeaderHandle(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintf(w, "The Web has not finished!")
}

func RedirectHandle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "www.google.com")
	w.WriteHeader(302)
}

func JsonHandle(w http.ResponseWriter, r *http.Request) {
	p := Post{
		User:    "Matthew",
		Threads: []string{"First", "Second", "Third"},
	}
	b, err := json.Marshal(p)
	if err == nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write(b)
	}
}

func main() {
	http.HandleFunc("/", WelcomeHandle)
	http.HandleFunc("/request", RequestHandle)
	http.HandleFunc("/body", BodyHandle)
	http.HandleFunc("/write", WriteHandle)
	http.HandleFunc("/writeheader", WriteHeaderHandle)
	http.HandleFunc("/redirect", RedirectHandle)
	http.HandleFunc("/json", JsonHandle)
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	server.ListenAndServe()
}
