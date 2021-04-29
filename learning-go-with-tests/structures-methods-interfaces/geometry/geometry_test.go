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

	verifyArea := func(t *testing.T, shape Shape, expected float64) {

		t.Helper()
		result := shape.Area()

		if result != expected {
			t.Errorf("result: %.2f, expected: %.2f", result, expected)
		}
	}

	t.Run("calc rectangler", func(t *testing.T) {

		rectangle := Rectangle{5.0, 5.0}
		verifyArea(t, rectangle, 25.0)
	})

	t.Run("calc circle", func(t *testing.T) {

		circle := Circle{10}
		verifyArea(t, circle, 314.1592653589793)
	})

	t.Run("calc triangle", func(t *testing.T) {

		triangle := Triangle{12, 6}
		verifyArea(t, triangle, 36.0)
	})
}
