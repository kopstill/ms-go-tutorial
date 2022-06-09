package test

import (
	"fmt"
	"testing"
)

func TestTest(t *testing.T) {
	// account := Account{"John", "Doe"}
	// fmt.Println(account)
	// account.ChangeName("Brad", "Pitt")

	// employee := Employee{account, 1000.00001}
	// fmt.Println(employee)
	// employee.AddCredits(100)
	// fmt.Println(employee)
	// employee.RemoveCredits(200)
	// fmt.Println(employee)
	// fmt.Println(employee.CheckCredits())

	bruce, _ := CreateEmployee("Bruce", "Lee", 500)
	fmt.Println(bruce.CheckCredits())
	credits, err := bruce.AddCredits(250)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("New Credits Balance =", credits)
	}

	_, err = bruce.RemoveCredits(2500)
	if err != nil {
		fmt.Println("Can't withdraw or overdrawn!", err)
	}

	bruce.ChangeName("Mark", "Wahlberg")

	fmt.Println(bruce)
}
