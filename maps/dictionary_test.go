package maps

import "testing"

func TestSearch(t *testing.T) {
	dict := map[string]string{"test": "this is just a test"}

	got := Search(dict, "test")
	expected := "this is just a test"

	assertStrings(t, got, expected)
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q want %q", got, expected)
	}
}
