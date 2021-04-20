package sum

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	result := Sum(5, 5)
	expected := 10

	if result != expected {
		t.Errorf("result: '%d', expected: '%d'", result, expected)
	}
}

func ExampleSum() {
	result := Sum(4, 4)
	fmt.Println(result)
	// Output: 8
}
