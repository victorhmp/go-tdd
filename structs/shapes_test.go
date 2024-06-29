package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	expected := 40.0

	if got != expected {
		t.Errorf("got %g expected %g", got, expected)
	}
}

func TestArea(t *testing.T) {
	areaTests := []struct {
		label    string
		shape    Shape
		expected float64
	}{
		{label: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, expected: 72.0},
		{label: "Circle", shape: Circle{Radius: 10}, expected: 314.1592653589793},
		{label: "Triangle", shape: Triangle{Base: 12, Height: 6}, expected: 36.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.label, func(t *testing.T) {
			got := tt.shape.Area()

			if got != tt.expected {
				t.Errorf("%#v got %g expected %g", tt.shape, got, tt.expected)
			}
		})
	}
}
