package structs

import "testing"

func TestPerimeter(t *testing.T) {

	perimeterTests := []struct {
		name         string
		shape        Shape
		hasPerimeter float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 12, Width: 6}, hasPerimeter: 36.0},
		{name: "Circle", shape: Circle{Radius: 9}, hasPerimeter: 56.548667764616276},
	}

	for _, tt := range perimeterTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Perimeter()
			if got != tt.hasPerimeter {
				t.Errorf("%#v got %g but wanted %g", tt.shape, got, tt.hasPerimeter)
			}
		})
	}
}

func TestArea(t *testing.T) {

	areaTests := []struct {
		name    string
		shape   Shape
		hasArea float64
	}{
		{name: "Rectangle", shape: Rectangle{Height: 12, Width: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 9}, hasArea: 254.46900494077323},
		{name: "Triangle", shape: Triangle{a: 12, b: 5, c: 13}, hasArea: 30.0},
	}

	for _, tt := range areaTests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area()
			if got != tt.hasArea {
				t.Errorf("%#v got %g but wanted %g", tt.shape, got, tt.hasArea)
			}
		})
	}
}
