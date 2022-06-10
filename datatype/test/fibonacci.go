package test

import (
	"errors"
	"fmt"
)

func fibonacci(num int) ([]int, error) {
	if num < 2 {
		return nil, errors.New("num should be greater than 2")
	}

	var result = make([]int, num)

	for i := 0; i < num; i++ {
		if i < 2 {
			result[i] = 1
		} else {
			result[i] = result[i-1] + result[i-2]
		}
	}

	return result, nil
}

func fibonacci1() {
	num := 15
	arr := make([]int, num)

	for i := 0; i < num; i++ {
		if i == 0 {
			arr[i] = 0
			fmt.Println(arr[i])
		} else if i == 1 {
			arr[i] = 1
			fmt.Println(arr[i])
		} else {
			arr[i] = arr[i-1] + arr[i-2]
			fmt.Println(arr[i])
		}
	}
}
