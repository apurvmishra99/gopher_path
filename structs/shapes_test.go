package structs

import "testing"

func TestPerimeter(t *testing.T) {
	rect := Rectangle{10.0, 10.0}
	got := Perimeter(rect)
	want := 40.0

	if got != want {
		t.Errorf("got %.2f but wanted %.2f", got, want)
	}
}

func TestArea(t *testing.T) {

	t.Run("calc area for rectangle", func(t *testing.T) {
		rect := Rectangle{12.0, 6.0}
		got := Area(rect)
		want := 72.0

		if got != want {
			t.Errorf("got %g but wanted %g", got, want)
		}
	})

	t.Run("calc area for circle", func(t *testing.T) {
		circ := Circle{9.0}
		got := Area(circ)
		want := 254.469004941

		if got != want {
			t.Errorf("got %g but wanted %g", got, want)
		}
	})

}
