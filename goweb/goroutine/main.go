package main

import (
	"fmt"
)

func printNum() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}
}

func printLetter() {
	for i := 'A'; i < 'A'+10; i++ {
		fmt.Println(i)
	}
}

func printTask1() {
	printNum()
	printLetter()
}

func printTask2() {
	go printNum()
	go printLetter()
}

func main() {

}
