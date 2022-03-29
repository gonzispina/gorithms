package fibonaccisumodds

// FibonacciSumOdds sums all the odd values that belongs to the fibonacci series that are less than a given value
// Solution in O(N)
func FibonacciSumOdds(N int) int {
	if N == 0 {
		return 0
	}

	if N == 1 {
		return 1
	}

	previous := make(map[int]int)
	previous[1] = 1
	previous[2] = 1

	sum := 2
	for i := 3; previous[i-1]+previous[i-2] <= N; i++ {
		fib := previous[i-1] + previous[i-2]
		previous[i] = fib
		if fib%2 != 0 {
			sum += fib
		}
	}

	return sum
}
