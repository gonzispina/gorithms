package countunorderedpairs

func count(a []int) ([]int, int) {
	if len(a) <= 1 {
		return []int{a[0]}, 0
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			return []int{a[1], a[0]}, 1
		}
		return []int{a[0], a[1]}, 0
	}

	mid := len(a) / 2
	left, c1 := count(a[0:mid])
	right, c2 := count(a[mid:])

	counter := c1 + c2

	i := 0
	j := 0
	res := make([]int, len(a))
	for c := 0; c < len(a); c++ {
		if i == len(left) {
			res[c] = right[j]
			j++
		} else if j == len(right) {
			res[c] = left[i]
			i++
		} else if left[i] < right[j] {
			res[c] = left[i]
			i++
		} else if left[i] > right[j] {
			res[c] = right[j]
			counter += len(left) - i
			j++
		} else if left[i] == right[j] {
			res[c] = left[i]
			i++
		}
	}

	return res, counter
}
