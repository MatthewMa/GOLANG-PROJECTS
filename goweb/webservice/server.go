package main

import (
	"encoding/json"
	"net/http"
	"path"
	"strconv"
)

type Post struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func main() {
	http.HandleFunc("/post/", HandleRequest)
	server := http.Server{
		Addr:    ":8080",
		Handler: nil,
	}
	server.ListenAndServe()
}

func HandleRequest(w http.ResponseWriter, r *http.Request) {
	var err error
	switch r.Method {
	case "GET":
		err = HandleGet(w, r)
	case "POST":
		err = HandlePost(w, r)
	case "PUT":
		err = HandlePut(w, r)
	case "DELETE":
		err = HandleDelete(w, r)
	}
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func HandleGet(w http.ResponseWriter, r *http.Request) (err error) {
	num, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil {
		return
	}
	post, err := retrieve(num)
	if err != nil {
		return
	}
	data, err := json.MarshalIndent(&post, "", "\t")
	if err != nil {
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
	return
}

func HandlePost(w http.ResponseWriter, r *http.Request) (err error) {
	len := r.ContentLength
	body := make([]byte, len)
	r.Body.Read(body)
	var post Post
	json.Unmarshal(body, &post)
	err = post.create()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func HandlePut(w http.ResponseWriter, r *http.Request) (err error) {
	var id int
	id, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	var post Post
	post, err = retrieve(id)
	if err != nil {
		return
	}
	len := r.ContentLength
	data := make([]byte, len)
	r.Body.Read(data)
	if err != nil {
		return
	}
	json.Unmarshal(data, &post)
	err = post.update()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}

func HandleDelete(w http.ResponseWriter, r *http.Request) (err error) {
	var num int
	num, err = strconv.Atoi(path.Base(r.URL.Path))
	if err != nil {
		return
	}
	post := Post{}
	post, err = retrieve(num)
	if err != nil {
		return
	}
	err = post.delete()
	if err != nil {
		return
	}
	w.WriteHeader(200)
	return
}
