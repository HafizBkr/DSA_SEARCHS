package ternarysearch


func TernarySearch(arr []int, target int) int {
	left, right := 0, len(arr)-1

	for left <= right {
		mid1 := left + (right-left)/3
		mid2 := left + 2*(right-left)/3

		if arr[mid1] == target {
			return mid1
		} else if arr[mid2] == target {
			return mid2
		} else if target < arr[mid1] {
			right = mid1 - 1
		} else if target > arr[mid2] {
			left = mid2 + 1
		} else {
			left, right = mid1+1, mid2-1
		}
	}

	return -1
}
