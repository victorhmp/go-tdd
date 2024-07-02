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

func TestAddWord(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		newWord := "test"
		newDefinition := "this is just a test"

		err := dictionary.Add(newWord, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, newWord, newDefinition)
	})

	t.Run("existing word", func(t *testing.T) {
		newWord := "test"
		newDefinition := "this is just a test"
		dictionary := Dictionary{newWord: newDefinition}

		err := dictionary.Add(newWord, newDefinition)

		assertError(t, err, ErrWordExists)
		assertDefinition(t, dictionary, newWord, newDefinition)
	})
}

func TestUpdateWord(t *testing.T) {
	t.Run("existing word", func(t *testing.T) {
		dictionary := Dictionary{"test": "this is just a test"}
		newDefinition := "new definition"
		word := "test"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, nil)
		assertDefinition(t, dictionary, word, newDefinition)
	})

	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		newDefinition := "this is just a test"

		err := dictionary.Update(word, newDefinition)

		assertError(t, err, ErrWordDoesNotExist)
	})
}

func TestDelete(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}
	word := "test"

	dictionary.Delete(word)

	_, searchError := dictionary.Search(word)
	assertError(t, searchError, ErrNotFound)
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)

	if err != nil {
		t.Fatal("should have found the word:", word)
	}

	assertStrings(t, got, definition)
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
