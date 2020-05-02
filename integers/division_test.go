package integers

import (
	"fmt"
	"testing"
)

func TestDivision(t *testing.T) {
	quotient, err := Division(10, 0)
	want := 3

	if err != nil {
		fmt.Println("Division failed due to: ", err)
		t.FailNow()
	}

	if quotient != want {
		t.Errorf("got '%d' but wanted '%d'", quotient, want)
	}

}
