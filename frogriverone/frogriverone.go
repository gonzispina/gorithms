package frogriverone

// FrogRiverOne tells which is the minimum time a frog has to wait to cross a river until all positions are covered
// Solution O(N)
func FrogRiverOne(x int, a []int) int {
	if x == 0 {
		return 0
	}

	if x < 0 || a == nil ||  len(a) == 0 || len(a) < x {
		return -1
	}

	if len(a) == x {
		return len(a) - 1
	}

	appearances := make([]bool, x)
	correctSum := x * (x + 1) / 2
	sum := 0
	res := -1
	for index, value := range a {
		if value > x || value <= 0 {
			continue
		}

		if !appearances[value-1] {
			appearances[value-1] = true
			sum += value
		}

		if sum == correctSum {
			res = index
			break
		}
	}

	return res
}
