package structs

import "math"

// Rectangle : defines the rectangle shape
type Rectangle struct {
	Height float64
	Width  float64
}

// Area : calculate the area of the rectangle
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

// Circle : defines the circle shape
type Circle struct {
	Radius float64
}

// Area : calculate the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Shape : implements the interface for Shapes
type Shape interface {
	Area() float64
}

// Perimeter : calculates the perimeter
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

// Area : calculates the area
func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}
