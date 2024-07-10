package interpolationsearch


func InterpolationSearch(arr []int, target int) int {
	low, high := 0, len(arr)-1

	for low <= high && target >= arr[low] && target <= arr[high] {
		pos := low + ((high - low) / (arr[high] - arr[low])) * (target - arr[low])

		if arr[pos] == target {
			return pos
		} else if arr[pos] < target {
			low = pos + 1
		} else {
			high = pos - 1
		}
	}

	return -1
}
