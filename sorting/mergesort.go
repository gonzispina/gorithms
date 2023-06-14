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
	a1 := MergeSort(a[0:mid])
	a2 := MergeSort(a[mid:])

	i := 0
	j := 0
	res := make([]int, len(a))
	for c := 0; c < len(a); c++ {
		if i == len(a1) {
			res[c] = a2[j]
			j++
			continue
		} else if j == len(a2) {
			res[c] = a1[i]
			i++
			continue
		}

		if a1[i] < a2[j] {
			res[c] = a1[i]
			i++
		} else if a1[i] > a2[j] {
			res[c] = a2[j]
			j++
		} else if a1[i] == a2[j] {
			res[c] = a1[i]
			i++
			c++
			res[c] = a2[j]
			j++
		}
	}

	return res
}
