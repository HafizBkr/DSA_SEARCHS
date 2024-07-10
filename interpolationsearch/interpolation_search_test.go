package interpolationsearch

import (
	"testing"
)

func TestInterpolationSearch(t *testing.T) {
	arr := []int{2, 3, 5, 8, 10, 12, 15}
	target := 8
	expected := 3

	result := InterpolationSearch(arr, target)
	if result != expected {
		t.Errorf("InterpolationSearch failed, got: %d, want: %d", result, expected)
	}

	targetNotFound := 7
	expectedNotFound := -1

	resultNotFound := InterpolationSearch(arr, targetNotFound)
	if resultNotFound != expectedNotFound {
		t.Errorf("InterpolationSearch failed for not found case, got: %d, want: %d", resultNotFound, expectedNotFound)
	}
}
