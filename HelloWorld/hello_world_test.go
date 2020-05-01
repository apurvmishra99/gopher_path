package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Apurv")
	want := "Hello, Apurv!"

	if got != want {
		t.Errorf("Got %q wanted %q", got, want)
	}
}
