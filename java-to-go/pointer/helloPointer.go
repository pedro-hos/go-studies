package main

import "fmt"

func main() {
	var str *string
	greeting := "hello world"
	str = &greeting

	fmt.Println("Address of Greeting:", str)
	fmt.Println("Greeting:", greeting)
	fmt.Println("Value stored in str:", *str)
}
