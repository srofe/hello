package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("The 'Hello' function shall return a message using the string passed.", func(t *testing.T) {
		got := Hello("Simon")
		want := "Hello, Simon"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

	t.Run("The 'Hello' function shall return a default message if empty string passed.", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
