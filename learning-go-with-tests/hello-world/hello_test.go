package hello_world

import "testing"

func TestHello(t *testing.T) {
	result := Hello()
	expected := "Hello, World!"

	if result != expected {
		t.Errorf("result: '%s', expected: '%s'", result, expected)
	}
}
