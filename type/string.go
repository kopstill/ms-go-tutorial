package main

import "fmt"

func main() {
	var firstName string = "John"
	lastName := "Doe"
	fmt.Println(firstName, lastName)

	fullName := "John Doe \t(alias \"Foo\")\n"
	fmt.Println(fullName)
}
