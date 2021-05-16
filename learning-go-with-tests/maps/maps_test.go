package maps

import "testing"

func TestSearch(t *testing.T) {

	m := map[string]string{"test": "just a test!"}

	got := Search(m, "test")
	expected := "just a test!"

	if got != expected {
		t.Errorf("got: '%s', expected: '%s'", got, expected)
	}
}
