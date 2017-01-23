package main

import (
	"encoding/json"
	"fmt"
	"io"
	_ "io/ioutil"
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

func main() {
	file, err := os.Open("person.json")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer file.Close()
	/*data := make([]byte, 0)
	data, err = ioutil.ReadAll(file)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	var post Post
	err = json.Unmarshal(data, &post)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println(post)*/
	var post Post
	decoder := json.NewDecoder(file)
	for {
		err := decoder.Decode(&post)
		if err == io.EOF {
			break
		}
		if err != nil {
			fmt.Println("Error:", err)
			os.Exit(1)
		}
		fmt.Println(post)
	}
}
