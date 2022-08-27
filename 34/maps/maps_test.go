package maps

import "testing"

func TestSearch(t *testing.T) {
	d := Dictionary{}
	t.Run("return definition of known word", func(t *testing.T) {
		word := "test"
		definition := "the means by which the presence, quality, or genuineness of anything is determined; a means of trial"
		err := d.Add(word, definition)
		assertErrorType(t, err, nil)
		assertDefinition(t, d, word, definition)
	})
	t.Run("return useful message when word not found", func(t *testing.T) {
		_, err := d.Search("unknown")
		if err == nil {
			t.Fatal("expected an error and did not receive one")
		}
		assertErrorType(t, err, ErrNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		word := "add"
		definition := "to unite or join so as to increase the number, quantity, size, or importance"
		err := d.Add(word, definition)
		assertErrorType(t, err, nil)
		assertDefinition(t, d, word, definition)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "add"
		definition := "to unite or join so as to increase the number, quantity, size, or importance"
		d := Dictionary{word: definition}
		err := d.Add(word, definition)
		assertErrorType(t, err, ErrAddExistingWord)
		assertDefinition(t, d, word, definition)
	})
}

func TestUpdate(t *testing.T) {
	word := "update"
	definition := "to incorporate new or more accurate information in (a database, program, procedure, etc.)"
	newDefinition := "to bring (a book, figures, or the like) up to date as by adding new information or making corrections"
	t.Run("existing word", func(t *testing.T) {
		d := Dictionary{word: definition}
		err := d.Update(word, newDefinition)
		assertErrorType(t, err, nil)
		assertDefinition(t, d, word, newDefinition)
	})
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		err := d.Update(word, newDefinition)
		assertErrorType(t, err, ErrUpdateExistingWord)
	})
}

func TestDelete(t *testing.T) {
	word := "delete"
	definition := "to strike out or remove (something written or printed); cancel; erase; expunge"
	t.Run("existing word", func(t *testing.T) {
		d := Dictionary{word: definition}
		err := d.Delete(word)
		assertErrorType(t, err, nil)
		_, err = d.Search(word)
		assertErrorType(t, err, ErrNotFound)
	})
	t.Run("new word", func(t *testing.T) {
		d := Dictionary{}
		err := d.Delete(word)
		assertErrorType(t, err, ErrDeleteNotFound)
	})
}

func assertDefinition(t testing.TB, d Dictionary, word, definition string) {
	t.Helper()
	got, err := d.Search(word)
	if err != nil {
		t.Fatal("should find the added word:", err)
	}
	if got != definition {
		t.Errorf("got definition: %q want definition: %q", got, definition)
	}
}

func assertErrorType(t testing.TB, gotErr, wantErr error) {
	t.Helper()
	if gotErr != wantErr {
		t.Errorf("expected the error %q but got %q", gotErr, wantErr)
	}
}
