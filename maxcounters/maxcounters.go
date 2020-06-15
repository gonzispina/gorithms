package maxcounters

// MaxCounters returns an array of counters of length n which counts the number of appearances of any integer value in a
// if a[i] = n + 1 then all the counters must be equalled to the maximum counter that exists in that moment
// Solution in O(N + M)
func MaxCounters(n int, a[]int) []int {
	if n == 0 {
		return []int{}
	}

	res := make([]int, n)
	if a == nil || len(a) == 0 {
		return res
	}

	max := 0
	maximized := 0
	for _, value := range a {
		if value > n + 1 || value < 1 {
			continue
		}

		if value == n + 1 {
			maximized = max
			continue
		}

		if res[value-1] < maximized {
			res[value-1] = maximized
		}

		res[value-1]++
		if max < res[value-1] {
			max = res[value-1]
		}
	}

	for index, value := range res {
		if value < maximized {
			res[index] = maximized
		}
	}

	return res
}
