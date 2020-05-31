package palindromesum

// PalindromeSum returns the sum of all the palindromes that are smaller than a given value
// Solution in O(N^2)
func PalindromeSum(n int) int {
	sum := 0
	v := 0
	for i := 1;; i++ {
		a := i
		v = i

		for a > 0 {
			v = v * 10 + (a % 10)
			a = a / 10
		}

		if v > n {
			break
		}

		sum += v
	}

	v = 0
	for i := 1;; i++ {
		a := i / 10
		v = i

		for a > 0 {
			v = v * 10 + (a % 10)
			a = a / 10
		}

		if v > n {
			break
		}

		sum += v
	}

	return sum
}
