package errorlog

import (
	"errors"
	"fmt"
	"time"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var ErrNotFound = errors.New("Employee not found")

func getInformation(id int) (*Employee, error) {
	for tries := 0; tries < 3; tries++ {
		employee, err := apiCallEmployee(id)
		if err == nil {
			return employee, nil
		}
		fmt.Println("Server is not responding, retrying ...")
		time.Sleep(time.Second * 2)
	}

	return nil, fmt.Errorf("server has failed to respond to get the employee information")
}

func apiCallEmployee(id int) (*Employee, error) {
	if id != 1001 {
		return nil, ErrNotFound
	}

	employee := Employee{ID: id, LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
