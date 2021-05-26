package maps

import "testing"

func compareString(t *testing.T, got string, expected string) {
	t.Helper()

	if got != expected {
		t.Errorf("got: '%s', expected: %s", got, expected)
	}
}

func TestSearch(t *testing.T) {
	dict := Dict{"test": "just a test!"}

	t.Run("search known", func(t *testing.T) {
		got, _ := dict.Search("test")
		expected := "just a test!"

		compareString(t, got, expected)
	})

	t.Run("search unknown", func(t *testing.T) {
		_, err := dict.Search("foo")

		if err == nil {
			t.Fatal("expected to get an error.")
		}
	})
}
