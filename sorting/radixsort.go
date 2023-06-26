package sorting

func RadixSort(arr []int) []int {
	max := findMax(arr)
	exp := 1
	for ; max > 0; max /= 10 {
		arr = CountSort(arr, exp)
		exp *= 10
	}
	return arr
}
