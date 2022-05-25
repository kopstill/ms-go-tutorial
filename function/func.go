package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	// number0, err0 := strconv.Atoi(os.Args[0])
	// number1, _ := strconv.Atoi(os.Args[1])
	// number2, _ := strconv.Atoi(os.Args[2])
	// fmt.Println(number0, err0)
	// fmt.Println("Sum:", number1+number2)

	fmt.Println("Sum:", sum1(os.Args[1], os.Args[2]))

	sum, mul := calc(os.Args[1], os.Args[2])
	fmt.Println("Sum:", sum, "Mul:", mul)

	firstName := "John"
	updateName(firstName)
	fmt.Println(firstName)
	updateName1(&firstName)
	fmt.Println(firstName)
}

func sum(number1 string, number2 string) int {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2
}

func sum1(number1 string, number2 string) (result int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	result = int1 + int2
	return
}

func calc(number1, number2 string) (int, int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	return int1 + int2, int1 * int2
}

func calc1(number1, number2 string) (sum, mul int) {
	int1, _ := strconv.Atoi(number1)
	int2, _ := strconv.Atoi(number2)
	sum = int1 + int2
	mul = int1 * int2
	return
}

func updateName(firstName string) {
	firstName = "David"
}

func updateName1(firstName *string) {
	*firstName = "David"
}
