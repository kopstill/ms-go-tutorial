package test

import "errors"

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
