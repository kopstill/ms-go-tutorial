package datatype

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	ID        int
	FirstName string `json:"first_name"`
	LastName  string
	Address   string `json:"address,omitempty"`
	Remark    string `json:"-"`
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
			ID:     1001,
			Remark: "memo",
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

	fmt.Println("---------------------------------")

	data, _ := json.Marshal(employee1)
	fmt.Println(data)
	fmt.Println(string(data))
	fmt.Printf("%s\n", data)

	var decoded Employee1
	json.Unmarshal(data, &decoded)
	fmt.Printf("%v\n", decoded)

	fmt.Println("---------------------------------")

	employees := []Employee1{
		Employee1{
			Person: Person{
				LastName: "Doe", FirstName: "John",
			},
			ManagerID: 1001,
		},
		Employee1{
			Person: Person{
				LastName: "Campbell", FirstName: "David",
			},
			ManagerID: 1002,
		},
	}

	data1, _ := json.Marshal(employees)
	fmt.Printf("%s\n", data1)

	var decoded1 []Employee1
	json.Unmarshal(data1, &decoded1)
	fmt.Printf("%v\n", decoded1)

	employeexxx := Employee{1001, "John", "Doe", "Doe's Street"} // won't warning in the same package without specifying the keys
	fmt.Println(employeexxx)
}
