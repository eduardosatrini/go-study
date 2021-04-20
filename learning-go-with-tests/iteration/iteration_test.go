package iteration

import (
	"testing"
)

func TestIterations(t *testing.T) {
	result := Iterations("d", 3)
	expected := "ddd"

	if result != expected {
		t.Errorf("result: %s, expected: %s", result, expected)
	}
}

func BenchmarkIterations(b *testing.B) {
	for i := 0; i < 5; i++ {
		Iterations("e", 60)
	}
}
