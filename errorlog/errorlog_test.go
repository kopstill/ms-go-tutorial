package errorlog

import (
	"fmt"
	"testing"
)

func TestEmployee(t *testing.T) {
	employee, err := getInformation(1001)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(employee)
	}
}
