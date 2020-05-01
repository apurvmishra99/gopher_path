package main

import "testing"

func TestHello(t *testing.T) {

	assertCorrectMessage := func(t *testing.T, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q wanted %q", got, want)
		}
	}

	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Apurv", "English")
		want := "Hello, Apurv!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'World'", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Spanish", func(t *testing.T) {
		got := Hello("Ale", "Spanish")
		want := "Hola, Ale!"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in Hindi", func(t *testing.T) {
		got := Hello("Ram", "Hindi")
		want := "Namaste, Ram!"
		assertCorrectMessage(t, got, want)
	})
}
