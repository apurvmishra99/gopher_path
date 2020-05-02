package integers

import "testing"

func TestMultiply(t *testing.T) {
	prod := Multiply(3, 3)
	want := 9

	if prod != want {
		t.Errorf("got '%d' but wanted '%d'", prod, want)
	}
}
