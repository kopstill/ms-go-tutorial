package test

import (
	"fmt"
	"testing"
)

func TestFibonacci(t *testing.T) {
	result, error := fibonacci(10)
	if error != nil {
		fmt.Println(error)
	}
	fmt.Println(result)
}

func TestRomanNumerals(t *testing.T) {
	fmt.Println("MCLX is:", romanToArabic("MCLX"))
	fmt.Println("MCMXCIX is:", romanToArabic("MCMXCIX"))
	fmt.Println("MCMZ is:", romanToArabic("MCMZ"))
}
