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
	t.Run("rectangles", func(t *testing.T) {
		rectangle := Rectangle{12, 6}
		got := rectangle.Area()
		expected := 72.0

		if got != expected {
			t.Errorf("got %.2f expected %.2f", got, expected)
		}
	})

	t.Run("circles", func(t *testing.T) {
		circle := Circle{10}
		got := circle.Area()
		expected := 314.1592653589793

		if got != expected {
			t.Errorf("got %g expected %g", got, expected)
		}
	})
}
