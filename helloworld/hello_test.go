package main

import "testing"

// There are some required bits here:
// - the test file has to end in `_test.go`
// - the test function only takes the one argument below
func TestHello(t *testing.T) {
	got := Hello("Victor")
	want := "Hello, Victor"

	if got != want {
		// `%q` wraps the value you're printing in double quotes
		t.Errorf("got %q want %q", got, want)
	}
}

// A super cool thing for Go is the `godoc` package. It enables you to launch
// a local server for a documentation portal with docs for all pkgs
// you have locally! Including all of the std library. So you can install
// it and check the docs for "testing" by going to http://localhost:<PORT>/pkg/testing/.
