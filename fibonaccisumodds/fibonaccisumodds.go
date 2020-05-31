package fibonaccisumodds

// FibonacciSumOdds sums all the odd values that belongs to the fibonacci series that are less than a given value
// Solution in O(N)
func FibonacciSumOdds(n int) int {
	if n == 0 {
		return 0
	}

	if n == 1 {
		return 1
	}

	t1:=0
	t2:=1
	nextTerm:=0
	sum := 1
	for i := 2; t1 + t2 <= n; i++ {
		nextTerm = t1 + t2
		t1 = t2
		t2 = nextTerm

		if t2 % 2 != 0 {
			sum += t2
		}
	}

	return sum
}