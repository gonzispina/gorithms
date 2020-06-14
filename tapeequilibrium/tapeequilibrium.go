package tapeequilibrium

// TapeEquilibrium splits the array in two parts at every index and calculates the difference between
// the sum of all the elements of the left and the sum of all the elements to the right and returns
// the smaller result
// Solved in O(N)
func TapeEquilibrium(a []int) int {
	if a == nil || len(a) == 0 {
		return 0
	}

	total := 0
	for _, value := range a {
		total += value
	}

	right := total - a[0]
	left := a[0]
	min := abs(right - left)

	for _, value := range a[1:] {
		if min == 0 {
			return 0
		}

		left += value
		right -= value
		diff := abs(right - left)

		if diff < min {
			min = diff
		}
	}

	return min
}

func abs(n int) int {
	if n < 0 {
		return -n
	}

	return n
}

