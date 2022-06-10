package generics

import "fmt"

func generics() {
	intarr := []int{12, -10, 5, 23, 42, -7, 16}
	fmt.Println(index(intarr, 42))

	strarr := []string{"hello", "world", "how", "are", "you"}
	fmt.Println(index(strarr, "ok"))
}

func index[T comparable](params []T, val T) int {
	for i, v := range params {
		if v == val {
			return i
		}
	}

	return -1
}
