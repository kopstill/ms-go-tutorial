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

func TestSimpleLog(t *testing.T) {
	simpleLog()
}

func TestPanicLog(t *testing.T) {
	panicLog()
}

func TestPrefixLog(t *testing.T) {
	prefixLog()
}

func TestFileLog(t *testing.T) {
	fileLog()
}
