// Package maps uses a map implementing a dictionary to illustrate testing, custom error handling and CRUD operations
package maps

const (
	ErrNotFound           = DictionaryErr("could not find that word in the dictionary")
	ErrDeleteNotFound     = DictionaryErr("cannot delete as word is not in the dictionary")
	ErrAddExistingWord    = DictionaryErr("cannot add as this word already exists in the dictionary")
	ErrUpdateExistingWord = DictionaryErr("cannot update as this word already exists in the dictionary")
)

// DictionaryErr allows attaching of Error and custom error types
type DictionaryErr string

// Error allows custom errors to be passed anywhere a standard library error is used
func (e DictionaryErr) Error() string {
	return string(e)
}

// Dictionary provides a custom type and name to our map
type Dictionary map[string]string

// Search finds the definition if the word is in dictionary or the ErrNotFound error if not in dictionary
func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

// Add places the word and definition in the dictionary returning the ErrAddExistingWord error if the word exists
func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		d[word] = definition
	case nil:
		return ErrAddExistingWord
	default:
		return err
	}
	return nil
}

// Update replaces the definition of the provided word, if available, and returns the ErrUpdateExistingWord error if not
func (d Dictionary) Update(word string, definition string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrUpdateExistingWord
	case nil:
		d[word] = definition
	default:
		return err
	}
	return nil
}

// Delete removes the provided word and definition, if available, and returns the ErrDeleteNotFound error if not
func (d Dictionary) Delete(word string) error {
	_, err := d.Search(word)
	switch err {
	case ErrNotFound:
		return ErrDeleteNotFound
	case nil:
		delete(d, word)
	default:
		return err
	}
	return nil
}
