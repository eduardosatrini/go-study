package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	result := Perimeter(5.0, 5.0)
	expected := 20.0

	if result != expected {
		t.Errorf("result: %.2f, expected: %.2f", result, expected)
	}

}

func TestArea(t *testing.T) {

	result := Area(5.0, 5.0)
	expected := 25.0

	if result != expected {
		t.Errorf("result: %.2f, expected: %.2f", result, expected)
	}

}
