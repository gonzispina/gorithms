package distinct

func Distinct(numbers []int) int {
	if numbers == nil {
		return 0
	}

	viewed := map[int]struct{}{}
	count := 0
	for _, n := range numbers {
		if _, ok := viewed[n]; !ok {
			viewed[n] = struct{}{}
			count++
		}
	}

	return count
}
