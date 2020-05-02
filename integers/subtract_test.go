package integers

import "testing"

func TestSubtract(t *testing.T) {
	res := Subtract(6, 3)
	want := 3

	if res != want {
		t.Errorf("got '%d' but wanted '%d'", res, want)
	}
}
