package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	}

	t.Run("The 'Hello' function shall return a message using the string passed.", func(t *testing.T) {
		got := Hello("Simon", "")
		want := "Hello, Simon"
		assertCorrectMessage(t, got, want)
	})

	t.Run("The 'Hello' function shall return a default message if empty string passed.", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("The 'Hello' function shall return a greeting in Spanish when that language is requested as an argument.", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("The 'Hello' function shall return a greeting in French when that language is requested as an argument.", func(t *testing.T) {
		got := Hello("Maxence", "French")
		want := "Bonjour, Maxence"
		assertCorrectMessage(t, got, want)
	})
}
