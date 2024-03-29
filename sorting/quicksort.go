package sorting

import "math/rand"

func QuickSort(a []int) []int {
	pending := [][]int{a}
	for len(pending) > 0 {
		arr := pending[0]
		pending = pending[1:]

		if len(arr) <= 1 {
			continue
		} else if len(arr) == 2 {
			if arr[0] > arr[1] {
				arr[0], arr[1] = arr[1], arr[0]
			}
			continue
		}

		pivot := rand.Intn(len(arr))
		arr[0], arr[pivot] = arr[pivot], arr[0]

		i := 0
		j := len(arr) - 1
		for i <= j {
			if arr[i] <= arr[0] {
				i++
				continue
			}

			if arr[j] >= arr[0] {
				j--
				continue
			}

			arr[i], arr[j] = arr[j], arr[i]
			i++
			j--
		}

		arr[0], arr[i-1] = arr[i-1], arr[0]

		pending = append(pending, arr[0:i])
		pending = append(pending, arr[i:])
	}

	return a
}
