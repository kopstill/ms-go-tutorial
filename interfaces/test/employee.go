package test

import (
	"fmt"
)

type Employee struct {
	Account
	credits float64
}

func (employee *Employee) AddCredits(credits float64) {
	employee.credits += credits
}

func (employee *Employee) RemoveCredits(credits float64) {
	employee.credits -= credits
}

func (employee Employee) CheckCredits() float64 {
	return employee.credits
}

func (employee Employee) String() string {
	// return employee.Account.String() + "\n" + strconv.FormatFloat(employee.credits, 'f', 2, 64)
	return employee.Account.String() + "\n" + fmt.Sprintf("%.2f", employee.credits)
	// return employee.Account.String() + "\n" + fmt.Sprintf("%v", employee.credits)
}
