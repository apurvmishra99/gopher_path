package integers

import "testing"

func TestDivision(t *testing.T) {
	quotient := Division(10, 3)
	want := 3

	if quotient != want {
		t.Errorf("got '%d' but wanted '%d'", quotient, want)
	}
}
