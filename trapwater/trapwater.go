package trapwater

func trap(height []int) int {
	var res int

	max := 0
	i := 0
	j := 0
	for i < len(height) {
		if height[i] == 0 {
			i++
			continue
		}

		if height[i] > max {
			max = height[i]
		}

		height[i]--

		hole := false
		for j = i + 1; j < len(height); j++ {
			if height[j] == 0 {
				hole = true
				continue
			}

			if height[j] > 0 {
				if hole {
					res += j - i - 1
					hole = false
				}
				i = j
			}

			if height[j] > max {
				max = height[j]
			}
			height[j]--
		}

		i++
	}

	if max > 1 {
		return res + trap(height)
	}

	return res
}
