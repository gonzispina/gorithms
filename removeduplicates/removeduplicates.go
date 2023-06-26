package removeduplicates

func removeDuplicates(nums []int) int {
	count := 1
	limit := len(nums)
	for i := 1; i < limit; {
		if nums[i-1] != nums[i] {
			count += 1
			i++
			continue
		}

		for j := i; j < len(nums)-1; j++ {
			nums[j] = nums[j+1]
		}

		limit -= 1
	}

	return count
}
