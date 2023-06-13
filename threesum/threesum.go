package threesum

import (
	"sort"
)

func threeSum(nums []int) [][]int {
	var res [][]int

	sort.Slice(nums, func(a, b int) bool {
		return nums[a] < nums[b]
	})

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i-1] == nums[i] {
			continue
		}

		low := i + 1
		high := len(nums) - 1

		sum := 0
		for low < high {
			sum = nums[i] + nums[low] + nums[high]
			if sum == 0 {
				res = append(res, []int{nums[i], nums[low], nums[high]})
				for low < high && nums[low+1] == nums[low] && nums[high-1] == nums[high] {
					high--
					low++
				}
				high--
				low++
			} else if sum > 0 {
				for low < high && nums[high-1] == nums[high] {
					high--
				}
				high--
			} else {
				for low < high && nums[low+1] == nums[low] {
					low++
				}
				low++
			}
		}
	}

	return res
}
