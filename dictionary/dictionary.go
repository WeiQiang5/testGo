package dictionary

import (
	"errors"
)

type Dictionary map[string]string
type Err string

func (e Err) Error() string {
	return string(e)
}

const (
	ErrNotFound         = Err("could not find the word you were looking for")
	ErrWordExists       = Err("current word is existing")
	ErrWordDoesNotExist = Err("cannot update word because it does not exist")
)

func (d Dictionary) Search(word string) (string, error) {
	definition, ok := d[word]
	if !ok {
		return "", ErrNotFound
	}
	return definition, nil
}

func (d Dictionary) Add(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		d[word] = definition
	case err == nil:
		return ErrWordExists
	default:
		return err
	}

	return nil
}

func (d Dictionary) Update(word, definition string) error {
	_, err := d.Search(word)

	switch {
	case errors.Is(err, ErrNotFound):
		return ErrWordDoesNotExist
	case err == nil:
		d[word] = definition
	default:
		return err
	}
	return err
}

func (d Dictionary) Delete(word string) {
	delete(d, word)
}
