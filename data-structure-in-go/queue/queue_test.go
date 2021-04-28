package queue

import (
	"testing"
)

func TestRemove(t *testing.T) {

	bankQueue := Queue{}
	bankQueue.Insert("John")
	bankQueue.Insert("Marie")
	bankQueue.Insert("Kate")
	bankQueue.Remove()

	result := bankQueue.values[0]
	expected := "Marie"

	if result != expected {
		t.Errorf("result: %s, expected: %s", result, expected)
	}

}
