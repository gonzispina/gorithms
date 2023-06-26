package minimumvsequence

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func findMinimum(arr []int) int {
	if len(arr) == 1 {
		return arr[0]
	}

	if len(arr) == 2 {
		return min(arr[0], arr[1])
	}

	low := 0
	high := len(arr) - 1

	for low < high {
		mid := (low + high) / 2

		if mid == 0 || mid == len(arr)-1 {
			return arr[mid]
		} else if arr[mid-1] < arr[mid] && arr[mid] < arr[mid+1] {
			high = mid
		} else if arr[mid-1] > arr[mid] && arr[mid] > arr[mid+1] {
			low = mid + 1
		} else {
			return arr[mid]
		}
	}

	return arr[low]
}
