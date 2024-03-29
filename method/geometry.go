package method

type triangle struct {
	size int
}

type square struct {
	size int
}

func (t triangle) perimeter() int {
	return t.size * 3
}

func (s square) perimeter() int {
	return s.size * 4
}

func (t *triangle) triangleDoubleSize() {
	t.size *= 2
}

func (s *square) squareDoubleSize() {
	s.size *= 2
}
