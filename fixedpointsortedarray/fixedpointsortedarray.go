package fixedpointsortedarray

func fixedPoint(arr []int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		m := (l + r) / 2
		if arr[m]-1 == m {
			return m + 1
		}

		if arr[m]-1 < m {
			l = m + 1
		} else {
			r = m
		}
	}

	if arr[l]-1 == l {
		return l + 1
	}

	return -1
}
