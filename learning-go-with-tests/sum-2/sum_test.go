package sum

import (
	"fmt"
	"testing"
)

type tableTests struct {
	desc    string
	numbers []int
	sum     int
}

func TestSum(t *testing.T) {
	t.Run("test cases sum", func(t *testing.T) {
		tt := []tableTests{
			{"sum five numbers", []int{1, 2, 3, 4, 5}, 15},
			{"sum with no values", []int{}, 0},
			{"sum with nil", nil, 0},
			{"sum with 1 + -1", []int{1, -1}, 0},
		}

		for _, tc := range tt {
			got := Sum(tc.numbers...)
			expected := tc.sum

			if got != expected {
				t.Errorf("desc: '%s', got: '%d', expected: '%d'", tc.desc, got, expected)
			}
		}

	})
}

func ExampleSum() {
	r := Sum(4, 5, 5)
	fmt.Println(r)
	// Output: 14
}
