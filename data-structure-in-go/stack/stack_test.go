package stack

import (
	"testing"
)

func TestPush(t *testing.T) {

	books := Stack{}
	result := books.Push("My German Brother")
	expected := 1

	if result != expected {
		t.Errorf("length stack - result: %d, expected: %d", result, expected)
	}

}

func TestPop(t *testing.T) {

	books := Stack{}
	books.Push("Dom Casmurro")
	books.Push("My German Brother")
	books.Push("Complete Histories")
	result := books.Pop()
	expected := 2

	if result != expected {
		t.Errorf("length stack - result: %d, expected: %d", result, expected)
	}

}
