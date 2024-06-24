package structs

import "testing"

func TestRectPerimeter(t *testing.T) {
	got := RectPerimeter(10.0, 10.0)
	expected := 40.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}

func TestRectArea(t *testing.T) {
  got := RectArea(12.0, 6.0)
	expected := 72.0

	if got != expected {
		t.Errorf("got %.2f expected %.2f", got, expected)
	}
}
