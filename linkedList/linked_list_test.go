package linkedList

import (
	"reflect"
	"testing"
)

func TestLinkedListInsertAndGet(t *testing.T) {
	ll := GetLinkedList()
	vals := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, val := range vals {
		ll.Insert(val)
	}

	actual := make([]int, len(vals))
	for i, _ := range vals {
		actual[i] = ll.Get(i)
	}

	ll.deleteNode(8)
	ll.print()

	if ll.Len() != len(vals) {
		t.Error(
			"For", vals,
			"expected", len(vals),
			"got", ll.Len(),
		)
	}
	if !reflect.DeepEqual(actual, vals) {
		t.Fatalf("Expected: %v, actual: %v\n", vals, actual)
	}
}

func TestLinkedListConstructor(t *testing.T) {
	vals := []int{1, 2, 3, 4}
	ll := GetLinkedListFromValues(vals)
	actual := make([]int, len(vals))

	for i, _ := range vals {
		actual[i] = ll.Get(i)
	}

	if ll.Len() != len(vals) {
		t.Error(
			"For", vals,
			"expected", len(vals),
			"got", ll.Len(),
		)
	}
	if !reflect.DeepEqual(actual, vals) {
		t.Fatalf("Expected: %v, actual: %v\n", vals, actual)
	}
}
