package main

import (
	"encoding/xml"
	"fmt"
	"io/ioutil"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id      string   `xml:"id,attr"`
	Content string   `xml:"content"`
	Author  Author   `xml:"author"`
}

type Author struct {
	Id   string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main() {
	post := Post{
		Id:      "1",
		Content: " Hello World!",
		Author: Author{
			Id:   "2",
			Name: "Sau Sheong",
		},
	}
	data, err := xml.MarshalIndent(&post, "", "	")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	err = ioutil.WriteFile("post.xml", []byte(xml.Header+string(data)), 0644)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
}
