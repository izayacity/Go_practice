package linkedList

import (
	"reflect"
	"testing"
)

func TestRemoveMiddleNode(t *testing.T) {
	vals := []int{1, 1, 1, 2, 3, 3, 4, 5, 5, 6}
	ll := GetLinkedListFromValues(vals)

	ll.RemoveMiddleNode(2)
	expected := []int{1, 1, 2, 3, 3, 4, 5, 5, 6}
	actual := ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 1: Expected: %v, actual: %v\n", expected, actual)
	}

	ll.RemoveMiddleNode(0)
	expected = []int{1, 2, 3, 3, 4, 5, 5, 6}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 2: Expected: %v, actual: %v\n", expected, actual)
	}

	ll.RemoveMiddleNode(ll.count - 1)
	expected = []int{1, 2, 3, 3, 4, 5, 5}
	actual = ll.Slice()

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 3: Expected: %v, actual: %v\n", expected, actual)
	}
}
