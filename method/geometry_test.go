package method

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	var triangle = triangle{size: 10}
	triangle.triangleDoubleSize()
	fmt.Println(triangle.perimeter())
}

func TestSquare(t *testing.T) {
	square := square{15}
	fmt.Println(square.perimeter())
}

func TestDoubleSize(t *testing.T) {
	triangle := triangle{3}
	square := square{4}
	triangle.triangleDoubleSize()
	square.squareDoubleSize()
	fmt.Println(triangle)
	fmt.Println(square)
	fmt.Println(triangle.perimeter())
	fmt.Println(square.perimeter())
}
