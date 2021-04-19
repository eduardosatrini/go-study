package hello_name

import "testing"

func TestHelloName(t *testing.T) {

	verifyIfMessageIsCorretly := func(t *testing.T, result, expected string) {

		t.Helper()
		if result != expected {
			t.Errorf("result: '%s', expected: '%s'", result, expected)
		}
	}

	t.Run("Say hello in portuguese", func(t *testing.T) {

		result := HelloName("Eduardo", "portuguese")
		expected := "Olá, Eduardo"

		verifyIfMessageIsCorretly(t, result, expected)
	})

	t.Run("Say hello in spanish", func(t *testing.T) {

		result := HelloName("Henrique", "spanish")
		expected := "Hola, Henrique"

		verifyIfMessageIsCorretly(t, result, expected)
	})

	t.Run("say 'Hello World' when string 'name' and 'idiom' is a empty", func(t *testing.T) {

		result := HelloName("", "")
		expected := "Hello, World"

		verifyIfMessageIsCorretly(t, result, expected)
	})

	t.Run("Say 'hello' in french", func(t *testing.T) {

		result := HelloName("Edward", "french")
		expected := "Bonjour, Edward"

		verifyIfMessageIsCorretly(t, result, expected)
	})

	t.Run("Say hello in english if the language is not recognized", func(t *testing.T) {
		result := HelloName("Maria", "dogelang")
		expected := "Hello, Maria"

		verifyIfMessageIsCorretly(t, result, expected)
	})

}
