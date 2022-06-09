package test

import (
	"errors"
	"fmt"
)

type Employee struct {
	Account
	Credits float64
}

func CreateEmployee(firstName, lastName string, credits float64) (*Employee, error) {
	return &Employee{Account{firstName, lastName}, credits}, nil
}

func (employee *Employee) AddCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		employee.Credits += amount
		return employee.Credits, nil
	}
	return 0.0, errors.New("invalid credit amount")
}

func (e *Employee) RemoveCredits(amount float64) (float64, error) {
	if amount > 0.0 {
		if amount <= e.Credits {
			e.Credits -= amount
			return e.Credits, nil
		}
		return 0.0, errors.New("you can't remove more credits than the account has")
	}
	return 0.0, errors.New("you can't remove negative numbers")
}

func (employee Employee) CheckCredits() float64 {
	return employee.Credits
}

func (employee Employee) String() string {
	// return employee.Account.String() + "\n" + strconv.FormatFloat(employee.credits, 'f', 2, 64)
	// return employee.Account.String() + "\n" + fmt.Sprintf("%.2f", employee.credits)
	// return employee.Account.String() + "\n" + fmt.Sprintf("%v", employee.credits)
	return fmt.Sprintf("Name: %s %s - Credits: %.2f\n", employee.FirstName, employee.LastName, employee.Credits)
}
