package main

import "testing"

func TestHello(t *testing.T) {

	assertMessage := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q want %q", got, want)
		}

	}

	t.Run("saying hello to people", func(*testing.T) {
		got := Hello("", english)
		want := englishHelloPrefix + "World" + "!"

		assertMessage(t, got, want)
	})

	t.Run("saying hello to Chris", func(*testing.T) {
		got := Hello("Chris", english)
		want := englishHelloPrefix + "Chris" + "!"
		assertMessage(t, got, want)
	})

	t.Run("in Spanish", func(*testing.T) {
		got := Hello("Elodie", spanish)
		want := spanishHelloPrefix + "Elodie" + "!"
		assertMessage(t, got, want)
	})

	t.Run("is French", func(*testing.T) {
		got := Hello("Marie", "French")
		want := frenchHelloPrefix + "Marie" + "!"
		assertMessage(t, got, want)
	})

}
