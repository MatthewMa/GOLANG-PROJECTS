package main

import (
	"fmt"
	"net/http"
	"time"
)

func CookieHandle(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "name",
		Value:    "Matthew",
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

func GetCookieHandle(w http.ResponseWriter, r *http.Request) {
	c1 := http.Cookie{
		Name:     "name",
		Value:    "Matthew",
		HttpOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
	}

	c2 := http.Cookie{
		Name:     "age",
		Value:    "27",
		HttpOnly: true,
		MaxAge:   -1,
		Expires:  time.Unix(1, 0),
	}
	//c := r.Header["Cookie"]
	c, err := r.Cookie("name")
	if err != nil {
		fmt.Fprintln(w, "Cannot get cookies!")
		return
	}
	http.SetCookie(w, &c1)
	http.SetCookie(w, &c2)
	fmt.Fprintln(w, c)
}

func main() {
	http.HandleFunc("/setcookies", CookieHandle)
	http.HandleFunc("/getcookies", GetCookieHandle)
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	server.ListenAndServe()

}
