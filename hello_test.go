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
		got := Hello("")
		want := englishHelloPrefix + "World" + "!"

		assertMessage(t, got, want)
	})

	t.Run("saying hello to Chris", func(*testing.T) {
		got := Hello("Chris")
		want := englishHelloPrefix + "Chris" + "!"
		assertMessage(t, got, want)
	})

}
