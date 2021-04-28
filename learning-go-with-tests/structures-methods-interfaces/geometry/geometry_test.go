package geometry

import (
	"testing"
)

func TestPerimeter(t *testing.T) {

	rectangle := Rectangle{5.0, 5.0}
	result := rectangle.Perimeter()
	expected := 20.0

	if result != expected {
		t.Errorf("result: %.2f, expected: %.2f", result, expected)
	}

}

func TestArea(t *testing.T) {

	t.Run("calc rectangler", func(t *testing.T) {
		rectangle := Rectangle{5.0, 5.0}
		result := rectangle.Area()
		expected := 25.0

		if result != expected {
			t.Errorf("result: %.2f, expected: %.2f", result, expected)
		}
	})

	t.Run("calc circle", func(t *testing.T) {
		circle := Circle{10}
		result := circle.Area()
		expected := 314.1592653589793

		if result != expected {
			t.Errorf("result: %.2f, expected: %.2f", result, expected)
		}
	})

}
