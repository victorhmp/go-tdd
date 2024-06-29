package structs

import "math"

// Very important to notice here that we're not telling Go that Rectangle or
// Circle implement this interface. They both have a method called `Area` that
// has the exact same signature and returns the same float64 we're declaring here.
// So Go will infer that they are Shapes.
// "In Go interface resolution is implicit. If the type you pass in matches what the interface is asking for, it will compile."
// That's why our tests also work!
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width  float64
	Height float64
}

type Circle struct {
	Radius float64
}

type Triangle struct {
	Base   float64
	Height float64
}

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (t Triangle) Area() float64 {
	return 0.5 * t.Base * t.Height
}
