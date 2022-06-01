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
