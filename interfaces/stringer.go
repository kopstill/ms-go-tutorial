package interfaces

import "fmt"

type Person struct {
	Name, Country string
}

func (p Person) String() string {
	// return fmt.Sprintf("%s is from %s", p.Name, p.Country)
	return fmt.Sprintf("%v is from %v", p.Name, p.Country)
}
