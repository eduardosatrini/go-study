package hello_name

import "testing"

func TestHelloName(t *testing.T) {
	result := HelloName("Eduardo")
	expected := "Hello, Eduardo!"

	if result != expected {
		t.Errorf("result: '%s', expected: '%s'", result, expected)
	}
}
