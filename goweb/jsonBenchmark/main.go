package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
)

type Post struct {
	Id       int       `json:"id"`
	Content  string    `json:"content"`
	Author   Author    `json:"author"`
	Comments []Comment `json:"comments"`
}

type Author struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Comment struct {
	Id      int    `json:"id"`
	Content string `json:"content"`
	Author  string `json:"author"`
}

func Unmarshall(name string) (post Post, err error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}
	defer file.Close()
	data, err := ioutil.ReadAll(file)
	if err != nil {
		return
	}
	err = json.Unmarshal(data, &post)
	if err != nil {
		return
	}
	return
}

func Decode(name string) (post Post, err error) {
	file, err := os.Open(name)
	if err != nil {
		return
	}
	defer file.Close()
	decoder := json.NewDecoder(file)
	err = decoder.Decode(&post)
	if err != nil {
		return
	}
	return
}

func main() {
	fileName := "post.json"
	p1, err := Unmarshall(fileName)
	if err != nil {
		log.Println("Cannot unmarshall:", err.Error())
		fmt.Println("Cannot unmarshall:", err.Error())
		panic(err)
	}
	fmt.Println(p1)
	p2, err := Decode(fileName)
	if err != nil {
		log.Println("Cannot unmarshall:", err.Error())
		fmt.Println("Cannot unmarshall:", err.Error())
		panic(err)
	}
	fmt.Println(p2)
}
