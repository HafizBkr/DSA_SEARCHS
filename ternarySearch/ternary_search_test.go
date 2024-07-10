package ternarysearch

import (
	"testing"
)

func TestTernarySearch(t *testing.T) {
	arr := []int{2, 3, 5, 8, 10, 12, 15}
	target := 8
	expected := 3

	result :=TernarySearch(arr, target)
	if result != expected {
		t.Errorf("TernarySearch failed, got: %d, want: %d", result, expected)
	}
	targetNotFound := 7
	expectedNotFound := -1

	resultNotFound := TernarySearch(arr, targetNotFound)
	if resultNotFound != expectedNotFound {
		t.Errorf("TernarySearch failed for not found case, got: %d, want: %d", resultNotFound, expectedNotFound)
	}
}
