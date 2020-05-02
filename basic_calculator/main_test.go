package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got int, want int) {
		t.Helper()
		if got != want {
			t.Errorf("got '%d' but wanted '%d'", got, want)
		}
	}

	t.Run("testing addition, adding 3 and 3", func(t *testing.T) {
		got := calculator("+", 3, 3)
		want := 6
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing subtraction, subtracting 100 and 3", func(t *testing.T) {
		got := calculator("-", 100, 3)
		want := 97
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing multiplication, mulitplying 90 and 50", func(t *testing.T) {
		got := calculator("*", 90, 50)
		want := 4500
		assertCorrectMessage(t, got, want)
	})

	t.Run("testing divsion, divinding 100 and 10", func(t *testing.T) {
		got := calculator("/", 100, 10)
		want := 10
		assertCorrectMessage(t, got, want)
	})
}
