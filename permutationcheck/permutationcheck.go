package permutationcheck

// PermutationCheck tells if an array is a permutation or not
// Solution in O(N) time and O(N) space
func PermutationCheck(A []int) bool {
	if A == nil || len(A) == 0 {
		return false
	}

	n := len(A)
	res := n * (n + 1) / 2
	sum := 0

	appearances := make([]bool, len(A))
	for _, value := range A {
		if value <= 0 || value > n || appearances[value-1]{
			return false
		}

		appearances[value - 1] = true
		sum += value
	}

	return sum == res
}
