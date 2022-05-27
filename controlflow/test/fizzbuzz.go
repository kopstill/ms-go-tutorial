package test

import (
	"fmt"
)

func fizzbuzz() {
	for i := 1; i < 100; i++ {
		switch {
		// case i%3 == 0 && i%5 == 0:
		case i%15 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}
}
