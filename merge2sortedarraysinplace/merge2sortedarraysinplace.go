package merge2sortedarraysinplace

func merge(nums1 []int, m int, nums2 []int, n int) {
	ptr1 := len(nums1) - len(nums2) - 1
	ptr2 := len(nums2) - 1

	index := len(nums1) - 1
	for index >= 0 {
		if ptr2 < 0 {
			break
		} else if ptr1 < 0 {
			nums1[index] = nums2[ptr2]
			ptr2--
		} else if nums1[ptr1] < nums2[ptr2] {
			nums1[index] = nums2[ptr2]
			ptr2--
		} else if nums1[ptr1] > nums2[ptr2] {
			nums1[index] = nums1[ptr1]
			ptr1--
		} else if nums1[ptr1] == nums2[ptr2] {
			nums1[index] = nums2[ptr2]
			ptr2--
			index--
			nums1[index] = nums1[ptr1]
			ptr1--
		}

		index--
	}
}
