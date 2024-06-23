package main

import "testing"

// There are some required bits here:
// - the test file has to end in `_test.go`
// - the test function only takes the one argument below
func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Victor")
		want := "Hello, Victor"

		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	// This is just telling the test suite that this is helper method
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

// A super cool thing for Go is the `godoc` package. It enables you to launch
// a local server for a documentation portal with docs for all pkgs
// you have locally! Including all of the std library. So you can install
// it and check the docs for "testing" by going to http://localhost:<PORT>/pkg/testing/.
