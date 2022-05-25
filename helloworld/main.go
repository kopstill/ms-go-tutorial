package main

import (
	"fmt"

	"github.com/kopever/calculator"

	calc "kopstill/calculator"

	"rsc.io/quote"
)

func main() {
	fmt.Println("Hello World!")
	total := calculator.Sum(3, 5)
	fmt.Println(total)
	fmt.Println("Version:", calculator.Version)

	fmt.Println(calc.Sum(1, 5))
	fmt.Println(calc.Version)

	fmt.Println(quote.Hello())
}
