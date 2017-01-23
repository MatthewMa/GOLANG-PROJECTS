package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func process(w http.ResponseWriter, r *http.Request) {
	//法一：
	// r.ParseMultipartForm(1024)
	// file := r.MultipartForm.File["uploaded"][0]
	// data, err := file.Open()
	// if err == nil {
	// 	filedata, err := ioutil.ReadAll(data)
	// 	if err == nil {
	// 		fmt.Fprintln(w, string(filedata))
	// 	}
	// }
	//法二：
	file, _, err := r.FormFile("uploaded")
	if err == nil {
		data, err := ioutil.ReadAll(file)
		if err == nil {
			fmt.Fprintln(w, string(data))
		}
	}
}

func simple(w http.ResponseWriter, r *http.Request) {
	str := `<html><title>GO WEB PROGRAMMING</title><body><h1>Hello,world!</h1></body></html>`
	w.Write([]byte(str))
}

func header(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(501)
	fmt.Fprintln(w, "Not implemented!")
}

func redirect(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Location", "http://www.google.ca")
	w.WriteHeader(302)
}

func setcookies(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "name",
		Value:    "matthew",
		HttpOnly: true,
	}
	c2 := http.Cookie{
		Name:     "age",
		Value:    "27",
		HttpOnly: true,
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
}

func getcookies(w http.ResponseWriter, r *http.Request) {
	c := r.Header["Cookie"]
	fmt.Fprintln(w, c)
}
func main() {
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/simple", simple)
	http.HandleFunc("/header", header)
	http.HandleFunc("/redirect", redirect)
	http.HandleFunc("/setcookies", setcookies)
	http.HandleFunc("/getcookies", getcookies)
	server.ListenAndServe()
}
