package linearsearch

import (
	"testing"
)

func TestLinearSearch(t *testing.T) {
	arr := []int{2, 3, 5, 8, 10, 12, 15}
	target := 8
	expected := 3

	result := LinearSearch(arr, target)
	if result != expected {
		t.Errorf("LinearSearch failed, got: %d, want: %d", result, expected)
	}
	targetNotFound := 7
	expectedNotFound := -1

	resultNotFound := LinearSearch(arr, targetNotFound)
	if resultNotFound != expectedNotFound {
		t.Errorf("LinearSearch failed for not found case, got: %d, want: %d", resultNotFound, expectedNotFound)
	}
}
