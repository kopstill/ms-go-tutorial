package interfaces

import (
	"fmt"
	"testing"
)

func TestShape(t *testing.T) {
	// var s Shape = Square{3}
	// fmt.Printf("%T\n", s)
	// fmt.Println(s.Area())
	// fmt.Println(s.Perimeter())

	// fmt.Println("----------")

	// var c Shape = Circle{10}
	// fmt.Printf("%T\n", c)
	// fmt.Println(c.Area())
	// fmt.Println(c.Perimeter())

	var s Shape = Square{3}
	printInformation(s)

	c := Circle{10}
	printInformation(c)
}

func printInformation(s Shape) {
	fmt.Printf("%T\n", s)
	fmt.Println("Area:", s.Area())
	fmt.Println("Perimeter:", s.Perimeter())
}

func TestStringer(t *testing.T) {
	rs := Person{"John Doe", "USA"}
	ab := Person{"Mark Collins", "United Kingdom"}
	fmt.Printf("%s\n%s\n", rs, ab)
	// fmt.Println(rs)
	// fmt.Println(ab)
}
