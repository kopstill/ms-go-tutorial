package method

import (
	"fmt"
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
