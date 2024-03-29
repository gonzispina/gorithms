package sorting

type SortingFunc func(a, b int) bool

func MergeSort(a []int) []int {
	if len(a) <= 1 {
		return []int{a[0]}
	}

	if len(a) == 2 {
		if a[0] > a[1] {
			return []int{a[1], a[0]}
		}
		return []int{a[0], a[1]}
	}

	mid := len(a) / 2
	left := MergeSort(a[0:mid])
	right := MergeSort(a[mid:])

	i := 0
	j := 0
	res := make([]int, len(a))
	for c := 0; c < len(a); c++ {
		if i == len(left) {
			res[c] = right[j]
			j++
			continue
		} else if j == len(right) {
			res[c] = left[i]
			i++
			continue
		}

		if left[i] < right[j] {
			res[c] = left[i]
			i++
		} else if left[i] > right[j] {
			res[c] = right[j]
			j++
		} else if left[i] == right[j] {
			res[c] = left[i]
			i++
			c++
			res[c] = right[j]
			j++
		}
	}

	return res
}
