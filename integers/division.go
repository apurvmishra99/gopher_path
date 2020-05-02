package integers

import (
	"errors"
)

// Division : Divide two integers and return their integral quotient
func Division(x, y int) (int, error) {
	if y == 0 {
		return 0, errors.New("Zero division error")
	}
	return x / y, nil
}
