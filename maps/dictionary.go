package main

type Dictionary map[string]string

type DictionaryError string

func (e DictionaryError) Error() string {
	return string(e)
}

const (
	ErrNotFound         = DictionaryError("could not find the word")
	ErrWordExists       = DictionaryError("key word already in dictionary")
	ErrWordDoesNotExits = DictionaryError("tried update not existing key")
)

func (d Dictionary) Search(word string) (string, error) {
	value, ok := d[word]
	if !ok {
		return value, ErrNotFound
	}
	return value, nil
}

func (d Dictionary) Add(key, value string) error {
	_, err := d.Search(key)
	switch err {
	case ErrNotFound:
		d[key] = value
	case nil:
		return ErrWordExists
	default:
		return err
	}
	return nil
}

func (d Dictionary) Update(key, newValue string) error {
	_, err := d.Search(key)
	if err == ErrNotFound {
		return ErrWordDoesNotExits
	}
	d[key] = newValue
	return nil
}

func (d Dictionary) Delete(key string) error {
	_, err := d.Search(key)
	switch err {
	case nil:
		delete(d, key)

	case ErrNotFound:
		return ErrWordDoesNotExits

	default:
		return err
	}
	return nil
}
