package hello_name

import "testing"

func TestHelloName(t *testing.T) {

	verifyIfMessageIsCorretly := func(t *testing.T, result, expected string) {
		t.Helper()
		if result != expected {
			t.Errorf("result: '%s', expected: '%s'", result, expected)
		}
	}

	t.Run("Say hello for users", func(t *testing.T) {
		result := HelloName("Eduardo")
		expected := "Hello, Eduardo!"

		verifyIfMessageIsCorretly(t, result, expected)
	})

	t.Run("say 'Hello World!' when string 'name' is a empty", func(t *testing.T) {
		result := HelloName("")
		expected := "Hello, World!"

		verifyIfMessageIsCorretly(t, result, expected)
	})

}
