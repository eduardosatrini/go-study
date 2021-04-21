package array

import (
	"fmt"
	"reflect"
	"testing"
)

func TestSum(t *testing.T) {

	slice := []int{5, 1, 2}
	result := Sum(slice)
	expected := 8

	if result != expected {
		t.Errorf("result: %v, expected: %v, value: %v", result, expected, slice)
	}

}

func ExampleSum() {

	result := Sum([]int{2, 2, 1})
	fmt.Println(result)
	// Output: 5

}

func TestSumAllValues(t *testing.T) {

	slice1 := []int{3, 2, 1, 4} // 10
	slice2 := []int{5, 2}       // 7
	result := SumAllValues(slice1, slice2)
	expected := []int{10, 7}

	if !reflect.DeepEqual(result, expected) {
		t.Errorf("result: %v, expected: %v", result, expected)
	}

}

func TestSumAllSlicesValues(t *testing.T) {

	slice1 := []int{2, 3, 0}    // 5
	slice2 := []int{1, 3}       // 4
	slice3 := []int{1, 1, 2, 1} // 5
	result := SumAllSlicesValues(slice1, slice2, slice3)
	expected := 14

	if result != expected {
		t.Errorf("result: %v, expected: %v", result, expected)
	}

}
