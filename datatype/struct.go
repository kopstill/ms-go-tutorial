package datatype

import "fmt"

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

var john Employee

func employee() {
	fmt.Println(john)
	fmt.Println(john == Employee{})

	employee := Employee{1001, "John", "Doe", "Doe's Street"}
	fmt.Println(employee)

	employee = Employee{LastName: "Doe", FirstName: "John"}
	fmt.Println(employee)
	employee.ID = 1002
	fmt.Println(employee)

	employeeCopy := &employee
	employee.FirstName = "David"
	fmt.Println(employeeCopy)
	fmt.Println(employee)
}
