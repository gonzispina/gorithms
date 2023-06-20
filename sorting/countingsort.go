package sorting

func CountSort(nums []int, exp int) []int {
	counts := make([]int, 10)
	for _, n := range nums {
		r := n / exp % 10
		counts[r]++
	}

	for i := 1; i < len(counts); i++ {
		counts[i] += counts[i-1]
	}

	output := make([]int, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		output[counts[(nums[i]/exp)%10]-1] = nums[i]
		counts[(nums[i]/exp)%10]--
	}

	return output
}
