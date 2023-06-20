package sorting

import (
	"fmt"
	"sort"
)

type TwoSlices struct {
	main  []int
	other []int
}

func (sbo *TwoSlices) Len() int {
	return len(sbo.main)
}

func (sbo *TwoSlices) Swap(i, j int) {
	sbo.main[i], sbo.main[j] = sbo.main[j], sbo.main[i]
	sbo.other[i], sbo.other[j] = sbo.other[j], sbo.other[i]
}

func (sbo *TwoSlices) Less(i, j int) bool {
	return sbo.main[i] > sbo.main[j]
}

func findMax(nums []int) int {
	max := nums[0]
	for _, n := range nums {
		if n > max {
			max = n
		}
	}
	return max
}

func largestNumber(nums []int) string {
	max := findMax(nums)

	exp := 1
	for max >= 10 {
		max = max / 10
		exp *= 10
	}

	exponents := make([]int, len(nums))
	for i, n := range nums {
		e := exp
		for n >= 10 {
			n = n / e
			e /= 10
		}
		exponents[i] = e
		nums[i] *= e
	}

	sort.Sort(&TwoSlices{
		main:  nums,
		other: exponents,
	})

	res := fmt.Sprintf("%v", nums[0]/exponents[0])
	for i := 1; i < len(nums); i++ {
		res = fmt.Sprintf("%s%v", res, nums[i]/exponents[i])
	}

	return res
}
