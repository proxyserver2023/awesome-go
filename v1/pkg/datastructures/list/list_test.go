package list

import "testing"

func TestList(t *testing.T) {
	t.Run("creating a basic list", func(t *testing.T) {
		got := ""
		want := ""
		if got != want {
			t.Errorf("GOT: `%s`. Expected: `%s`", got, want)
		}
	})
}
