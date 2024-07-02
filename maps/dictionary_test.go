package maps

import "testing"

func TestSearch(t *testing.T) {
	t.Run("known word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		got, _ := dictionary.Search("test")
		expected := "this is just a test"

		assertStrings(t, got, expected)
	})

	t.Run("unknown word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}

		_, err := dictionary.Search("unknown")

		if err == nil {
			t.Fatal("expected to get an error.")
		}

		assertError(t, err, ErrNotFound)
	})
}

func assertStrings(t testing.TB, got, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q want %q", got, expected)
	}
}

func assertError(t testing.TB, got, expected error) {
	t.Helper()

	if got != expected {
		t.Errorf("got %q, expected %q", got, expected)
	}
}
