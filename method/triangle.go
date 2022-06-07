package method

type coloredTriangle struct {
	// t     triangle
	triangle
	color string
}

func (ct coloredTriangle) perimeter() int {
	// return ct.triangle.perimeter() * 2
	return ct.size * 3 * 2
}
