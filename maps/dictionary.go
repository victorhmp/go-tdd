package maps

import "errors"

type Dictionary map[string]string

var ErrNotFound = errors.New("could not find the word you were looking for")

func (d Dictionary) Search(word string) (string, error) {
	// The second return value of a map access is a boolean value that reports
	// whether the requested key was present in the map or not.
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}
