package maps

type Dictionary map[string]string
type DictionaryErr string

const (
	ErrNotFound         = DictionaryErr("could not find the word you were looking for")
	ErrWordExists       = DictionaryErr("word already exists in the dictionary")
	ErrWordDoesNotExist = DictionaryErr("word does not exist in the dictionary")
)

func (e DictionaryErr) Error() string {
	return string(e)
}

func (d Dictionary) Search(word string) (string, error) {
	// The second return value of a map access is a boolean value that reports
	// whether the requested key was present in the map or not.
	definition, ok := d[word]

	if !ok {
		return "", ErrNotFound
	}

	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, newDefinition string) error {
	_, err := d.Search(word)

	switch err {
	case ErrNotFound:
		return ErrWordDoesNotExist
	case nil:
		d[word] = newDefinition
	default:
		return err
	}

	return nil
}

// Important to note about Maps: they are NOT reference types. A map value is a
// reference to a runtime hash table. So when you pass a map to a function, you
// are copying it, but just the **pointer**, not the underlying data.
// Also, note that maps can be `nil`. You can read from a `nil` map, but you
// can't write to it. Because of this "interesting feature", it's a good idea
// to **never** initialize a `nil` map and always use `make` to create a new map
// or explicitly create an empty map. Both of these are still ideomatic Go and
// and actually create empty hash maps, ensuring you don't get runtime panics.
