package cache_test

import (
	"reflect"
	"testing"
	"github.com/izayacity/Go_practice/cache"
)

func TestLRUCache(t *testing.T) {
	lru := cache.GetLRUCache(2)
	lru.Put(1, 1)
	lru.Put(2, 2)

	expected := 1
	actual := lru.Get(1)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 1: Expected: %v, actual: %v\n", expected, actual)
	}

	lru.Put(3, 3)
	expected = -1
	actual = lru.Get(2)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 2: Expected: %v, actual: %v\n", expected, actual)
	}

	lru.Put(4, 4)
	expected = -1
	actual = lru.Get(1)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 3: Expected: %v, actual: %v\n", expected, actual)
	}

	expected = 3
	actual = lru.Get(3)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 4: Expected: %v, actual: %v\n", expected, actual)
	}

	expected = 4
	actual = lru.Get(4)

	if !reflect.DeepEqual(expected, actual) {
		t.Fatalf("Test 4: Expected: %v, actual: %v\n", expected, actual)
	}
}