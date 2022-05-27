package test

import "fmt"

func findprimes(num int) bool {
	if num == 1 {
		return false
	}

	for i := 2; i < num; i++ {
		if num%i == 0 {
			return false
		}
	}

	// if num > 1 {
	// 	return true
	// } else {
	// 	return false
	// }

	return true
}

func printprimes() {
	fmt.Println("Prime numbers less than 20:")
	for i := 1; i < 20; i++ {
		if findprimes(i) {
			fmt.Printf("%d\n", i)
		}
	}
}
