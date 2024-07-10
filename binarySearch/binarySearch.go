package binarysearch

import (
	"fmt"
)

func BinarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1
	for left <= right {
		mid := (left+right)/2
		if arr[mid] == target {
			return mid
		} else if arr[mid] < target {
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return -1
}

func main() {
	arr := []int{2, 3, 40, 60, 100, 450} 
	target := 3
	result := BinarySearch(arr, target)
	if result != -1 {
		fmt.Println("Element found at index:", result)
	} else {
		fmt.Println("Element not found in array")
	}
}
