package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMethod := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("The 'Hello' function shall return a message using the string passed.", func(t *testing.T) {
		got := Hello("Simon")
		want := "Hello, Simon"
		assertCorrectMethod(t, got, want)
	})

	t.Run("The 'Hello' function shall return a default message if empty string passed.", func(t *testing.T) {
		got := Hello("")
		want := "Hello, world"
		assertCorrectMethod(t, got, want)
	})
}
