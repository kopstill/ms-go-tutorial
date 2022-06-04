package errorlog

import (
	"errors"
	"fmt"
	"testing"
)

func TestEmployee(t *testing.T) {
	employee, err := getInformation(1000)
	if err != nil {
		if errors.Is(err, ErrNotFound) {
			fmt.Printf("NOT FOUND: %v\n", err)
		} else {
			fmt.Println(err)
		}
	} else {
		fmt.Println(employee)
	}
}
