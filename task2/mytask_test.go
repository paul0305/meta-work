package task2

import "testing"

func TestShape(t *testing.T) {
	var s Shape
	s = Circle{2}
	s.Area()
	s.Perimeter()
	s = Rectangle{width: 3.25, height: 4.5}
	s.Area()
	s.Perimeter()
}
