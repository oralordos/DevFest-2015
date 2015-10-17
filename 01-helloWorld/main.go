package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	fmt.Println("What is your name?")
	var myName string
	fmt.Scanln(&myName)
	fmt.Println("It is nice to meet you,", myName)
}
