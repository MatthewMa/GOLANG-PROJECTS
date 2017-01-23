package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	data := []byte("Hello,world!\n")
	ioutil.WriteFile("myFile1", data, 0644)
	data1, err := ioutil.ReadFile("myFile1")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(data1))
	file, _ := os.Create("myFile2")
	file.Write(data)
	file1, err := os.Open("myFile2")
	defer file.Close()
	if err != nil {
		panic(err)
	}
	data2 := make([]byte, len(data))
	file1.Read(data2)
	defer file1.Close()
	fmt.Println(string(data2))
}
