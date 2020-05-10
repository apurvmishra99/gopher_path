package structs

// Rectangle : defines the rectangle shape
type Rectangle struct {
	Height float64
	Width  float64
}

// Circle : defines the circle shape
type Circle struct {
	Radius float64
}

// Perimeter : calculates the perimeter
func Perimeter(rect Rectangle) float64 {
	return 2 * (rect.Height + rect.Width)
}

// Area : calculates the area
func Area(rect Rectangle) float64 {
	return rect.Height * rect.Width
}
