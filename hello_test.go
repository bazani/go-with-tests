package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Bazani", "")
		want := "Hellow, Bazani"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hellow, Worldy' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hellow, Worldy"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Json", "French")
		want := "Bonjour, Json"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Portuguese", func(t *testing.T) {
		got := Hello("Ticia", "Portuguese")
		want := "Ola, Ticia"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
