package errorlog

type Employee struct {
	ID        int
	FirstName string
	LastName  string
	Address   string
}

func getInformation(id int) (*Employee, error) {
	employee, err := apiCallEmployee(id)
	return employee, err
}

func apiCallEmployee(id int) (*Employee, error) {
	employee := Employee{ID: id, LastName: "Doe", FirstName: "John"}
	return &employee, nil
}
