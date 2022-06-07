package method

import (
	"fmt"
	"kopever/ms-go-tutorial/datatype"
	"kopever/ms-go-tutorial/method/geometry"
	"testing"
)

func TestColoredTriangle(t *testing.T) {
	// triangle := coloredTriangle{triangle: triangle{5}, color: "blue"}
	triangle := coloredTriangle{triangle{5}, "blue"}
	// fmt.Println("Perimeter:", triangle.t.perimeter())
	fmt.Println("Color:", triangle.color)
	fmt.Println("Perimeter (colored):", triangle.perimeter())
	fmt.Println("Perimeter (normal):", triangle.triangle.perimeter())
}

func TestPackageTriangle(t *testing.T) {
	// triangle := geometry.Triangle{1} // warning: size not visiable
	triangle := geometry.Triangle{}
	triangle.SetSize(3)
	fmt.Println("Perimeter", triangle.Perimeter())

	// employee := datatype.Employee{1001, "John", "Doe", "Doe's Street"} // warning outside the definition package without specifying the keys
	employee := datatype.Employee{ID: 1001, FirstName: "John", LastName: "Doe", Address: "Doe's Street"}
	fmt.Println(employee)
}
