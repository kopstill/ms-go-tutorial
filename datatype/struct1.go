package datatype

import "fmt"

type Person struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

type Employee1 struct {
	// Information Person
	Person
	ManagerID int
}

type Contractor struct {
	Person
	CompanyID int
	FirstName string
	LastName  string
}

func employee1() {
	var employee Employee1
	// employee.Information.FirstName = "John"
	employee.FirstName = "John"
	fmt.Println(employee)

	employee1 := Employee1{
		Person: Person{
			ID: 1001,
		},
		ManagerID: 1002,
	}
	fmt.Println(employee1)
	fmt.Println(employee1.Person)
	fmt.Println(employee1.Person.FirstName)
	fmt.Println(employee1.LastName)
	fmt.Println(employee1.ManagerID)

	fmt.Println("---------------------------------")

	var contractor Contractor
	contractor.FirstName = "John"
	fmt.Println(contractor)
	contractor = Contractor{}
	fmt.Println(contractor)

	contractor = Contractor{
		Person: Person{
			ID:        1234,
			FirstName: "Hello",
			LastName:  "World",
			Address:   "Mars",
		},
		CompanyID: 5678,
	}
	contractor.FirstName = "John"
	contractor.LastName = "Wick"
	fmt.Println(contractor)
	contractor.Person.FirstName = "Tom"
	contractor.Person.LastName = "Hanks"
	fmt.Println(contractor)
}
