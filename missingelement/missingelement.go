package missingelement

// MissingElement returns the missing integer in an array
// Solution in O(N) based on that the sum of Ki when i=0 to n is equal to n * (n + 1) / 2
func MissingElement(a []int) int {
	if a == nil || len(a) == 0 {
		return 0
	}

	n := len(a) + 1
	correctSum := n * (n + 1) / 2

	wrongSum := 0
	for _, value := range a {
		wrongSum += value
	}

	return correctSum - wrongSum
}
