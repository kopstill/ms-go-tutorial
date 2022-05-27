package controlflow

import "fmt"

func recoverDemo() {
	defer func() {
		handler := recover()
		if handler != nil {
			fmt.Println("recoverDemo(): recover", handler)
		}
	}()

	highlow(2, 0)
	fmt.Println("Program finished successfully!")
}
