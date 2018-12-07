package gopl

import "testing"

func TestDuplicateWithStdInandFile(t *testing.T) {

	var assertEqual = func(t *testing.T, got, want string) {
		t.Helper()

		if got != want {
			t.Errorf("got '%s' want '%s'\n", got, want)
		}
	}

	t.Run("Printing Duplicates From StdLib", func(t *testing.T) {
		got := DuplicateTextStdInAndFile()
		want := ""
		assertEqual(t, got, want)
	})

}
