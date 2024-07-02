package maps

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	got := dictionary.Search("test")
	expected := "this is just a test"

	assertStrings(t, got, expected)
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q want %q", got, expected)
	}
}
