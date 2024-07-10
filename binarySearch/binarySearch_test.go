package binarysearch

import (
	"testing"
)

func TestBinarySearch(t *testing.T) {
	arr := []int{2, 3, 5, 8, 10, 12, 15}
	target := 8
	expected := 3

	result := BinarySearch(arr, target)
	if result != expected {
		t.Errorf("BinarySearch failed, got: %d, want: %d", result, expected)
	}
	targetNotFound := 7
	expectedNotFound := -1

	resultNotFound := BinarySearch(arr, targetNotFound)
	if resultNotFound != expectedNotFound {
		t.Errorf("BinarySearch failed for not found case, got: %d, want: %d", resultNotFound, expectedNotFound)
	}
}
