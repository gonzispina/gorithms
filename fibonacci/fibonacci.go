package fibonacci

func Fibonacci(N int) []int {
	if N == 0 {
		return []int{}
	}

	if N == 1 {
		return []int{0}
	}

	if N == 2 {
		return []int{0, 1}
	}

	res := make([]int, N)
	res[0] = 0
	res[1] = 1
	res[2] = 1

	for i := 3; i <= N-1; i++ {
		value := res[i-1] + res[i-2]
		res[i] = value
	}

	return res
}
