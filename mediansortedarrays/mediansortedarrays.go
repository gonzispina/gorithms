package mediansortedarrays

func calcMedian(arr []int) float64 {
	i := len(arr) / 2
	if len(arr)%2 == 1 {
		return float64(arr[i])
	} else {
		return float64(arr[i-1]+arr[i]) / 2
	}
}

func value(i int, a, b []int) int {
	if i < len(a) {
		return a[i]
	} else {
		return b[i-len(a)]
	}
}

func max(i int, i2 int) int {
	if i > i2 {
		return i
	}
	return i2
}

func min(i int, i2 int) int {
	if i < i2 {
		return i
	}
	return i2
}

func findMedianSortedArrays(a []int, b []int) float64 {
	if len(a) < len(b) {
		return findMedianSortedArrays(b, a)
	}

	if len(a) == 0 {
		return 0
	}

	i := len(a) / 2
	if len(a)%2 == 0 {
		i -= 1
	}

	median := calcMedian(a)
	if len(b) == 0 {
		return median
	}

	n := len(a)
	m := len(b)
	total := len(a) + len(b)
	middle := (total - 1) / 2
	if a[n-1] <= b[0] {
		if total%2 == 0 {
			return float64(value(middle, a, b)+value(middle+1, a, b)) / 2
		}
		return float64(value(middle, a, b))
	} else if b[m-1] <= a[0] {
		if total%2 == 0 {
			return float64(value(middle, b, a)+value(middle+1, b, a)) / 2
		}
		return float64(value(middle, b, a))
	}

	left := 0
	right := len(b) - 1
	for left+1 < right {
		med := (left + right) / 2

		if float64(b[med]) <= median {
			left = med
		} else {
			right = med
		}
	}

	i = int(float64(i) + float64(len(b[left+1:])-len(b[:left+1]))/2)

	if (len(a)+len(b))%2 == 1 {
		return float64(max(a[i], b[left]))
	}

	if len(b) == 1 {
		return float64(max(a[i], b[left])+a[i+1]) / 2
	}

	return float64(max(a[i], b[left])+min(a[i+1], b[left+1])) / 2
}
