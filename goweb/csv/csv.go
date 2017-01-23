package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

type Person struct {
	Name   string
	Age    int
	Gender string
}

func main() {
	file, _ := os.Create("myFile.csv")
	data := []Person{
		Person{"Matthew", 27, "M"},
		Person{"Joe", 26, "M"},
		Person{"Mary", 21, "F"},
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	for _, val := range data {
		line := []string{val.Name, strconv.Itoa(val.Age), val.Gender}
		err := writer.Write(line)
		if err != nil {
			panic(err)
		}
	}
	writer.Flush()
	file1, err := os.Open("myFile.csv")
	defer file1.Close()
	if err != nil {
		panic(err)
	}
	reader := csv.NewReader(file1)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}
	var persons []Person
	for _, item := range record {
		age, _ := strconv.ParseInt(item[1], 0, 0)
		person := Person{Age: int(age), Name: item[0], Gender: item[2]}
		persons = append(persons, person)
	}
	fmt.Println(persons[0].Name)
	fmt.Println(persons[0].Age)
	fmt.Println(persons[0].Gender)
}
