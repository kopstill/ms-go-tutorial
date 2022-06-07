package method

import (
	"fmt"
	"testing"
)

func TestTriangle(t *testing.T) {
	var triangle = triangle{size: 10}
	fmt.Println(triangle.perimeter())
}

func TestSquare(t *testing.T) {
	square := square{15}
	fmt.Println(square.perimeter())
}
