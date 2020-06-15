package missinginteger

// MissingInteger returns the minimum positive integer that is missing in a.
// N < 100.000 so given x as the missing integer then 1 <= x <= 100001
// Solution in O(N)
func MissingInteger(a []int) int {
	if a == nil || len(a) == 0 {
		return 1
	}

	appearances := make([]bool, 100000)
	for _, value := range a {
		if value <= 0 || value > 100000 {
			continue
		}

		appearances[value-1] = true
	}

	for index, value := range appearances {
		if !value {
			return index + 1
		}
	}

	return 100001
}
