package semiorderedlist

import "github.com/gonzispina/gorithms/sorting"

func sort(arr []int, low, high int) {
	var lower []int
	var middle []int
	var higher []int

	for _, v := range arr {
		if v <= low {
			lower = append(lower, v)
		} else if v >= high {
			higher = append(higher, v)
		} else {
			middle = append(middle, v)
		}
	}

	j := 0
	for i := len(lower) - 1; i >= 0; i-- {
		arr[j] = lower[i]
		j++
	}

	middle = sorting.MergeSort(middle) // O(sqrt(n)*log(sqrt(n)))
	for j = 0; j < len(middle); j++ {
		arr[j+len(middle)] = middle[j]
	}

	for j = 0; j < len(higher); j++ {
		arr[j+len(middle)+len(lower)] = higher[j]
	}
}
