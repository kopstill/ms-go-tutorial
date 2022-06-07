package test

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	account := Account{"John", "Doe"}
	fmt.Println(account)
	account.ChangeName("Brad", "Pitt")

	employee := Employee{account, 1000.00001}
	fmt.Println(employee)
	employee.AddCredits(100)
	fmt.Println(employee)
	employee.RemoveCredits(200)
	fmt.Println(employee)
	fmt.Println(employee.CheckCredits())
}
