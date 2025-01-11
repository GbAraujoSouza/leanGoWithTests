package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertString(t, got, want)
	})

	t.Run("unknown word", func(t *testing.T) {
		_, err := dictionary.Search("blue")
		want := ErrNotFound.Error()

		if err == nil {
			t.Fatal("expected to get an error")
		}

		assertString(t, err.Error(), want)
	})

}

func TestAdd(t *testing.T) {

	t.Run("new word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"

		dictionary := Dictionary{}
		err := dictionary.Add(key, value)

		assertError(t, err, nil)
		assertKeyValue(t, dictionary, key, value)
	})

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is the second test"

		dictionary := Dictionary{key: value}
		err := dictionary.Add(key, value)

		assertError(t, err, ErrWordExists)
		assertKeyValue(t, dictionary, key, value)
	})
}

func TestUpdate(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is the second test"
		dict := Dictionary{key: value}
		newValue := "new Value"

		err := dict.Update(key, newValue)

		assertError(t, err, nil)
		assertKeyValue(t, dict, key, newValue)
	})

	t.Run("new word", func(t *testing.T) {
		key := "test"
		dict := Dictionary{}
		newValue := "new Value"

		err := dict.Update(key, newValue)

		assertError(t, err, ErrWordDoesNotExits)
	})
}

func TestDelete(t *testing.T) {

	t.Run("existing word", func(t *testing.T) {
		key := "test"
		value := "this is just a test"
		dict := Dictionary{key: value}

		dict.Delete(key)
		_, err := dict.Search(key)

		assertError(t, err, ErrNotFound)
	})

	t.Run("word not in dictionary", func(t *testing.T) {
		key := "test"
		// value := "this is just a test"
		dict := Dictionary{}

    err := dict.Delete(key)

		assertError(t, err, ErrWordDoesNotExits)
	})
}

func assertKeyValue(t testing.TB, dict Dictionary, key string, value string) {
	t.Helper()
	got, err := dict.Search(key)
	if err != nil {
		t.Fatal("should found added word: ", err)
	}
	assertString(t, got, value)
}

func assertError(t testing.TB, err error, wantError error) {
	t.Helper()
	if err != wantError {
		t.Fatalf("got error %q want %q\n", err, wantError)
	}
}

func assertString(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}
