package array

import (
	"testing"
)

func TestSum(t *testing.T) {

	t.Run("sum slice integers", func(t *testing.T) {

		slice := []int{5, 1, 2}
		result := Sum(slice)
		expected := 8

		if result != expected {
			t.Errorf("result: %v, expected: %v, value: %v", result, expected, slice)
		}

	})

}
