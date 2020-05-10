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

// Perimeter : calculate the perimeter of the rectangle
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Height + r.Width)
}

// Circle : defines the circle shape
type Circle struct {
	Radius float64
}

// Area : calculate the area of the circle
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perimeter : calculate the perimeter of the circle
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// Triangle : defines the triangle shape
type Triangle struct {
	a float64
	b float64
	c float64
}

// Area : calculate the area of the triangle
func (t Triangle) Area() float64 {
	s := t.Perimeter() / 2
	area := math.Sqrt(s * ((s - t.a) * (s - t.b) * (s - t.c)))
	return area
}

// Perimeter : calculates the perimeter of the triangle
func (t Triangle) Perimeter() float64 {
	return t.a + t.b + t.c
}

// Shape : implements the interface for Shapes
type Shape interface {
	Area() float64
	Perimeter() float64
}

// Perimeter : calculates the perimeter
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

// Area : calculates the area
func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}
