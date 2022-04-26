package vendingmachine

func Solution(M float64, P float64) []int {
	res := make([]int, 6)
	if M < P {
		return res
	}

	coins := []int64{100, 50, 25, 10, 5, 1}
	m := int64(M * 100)
	p := int64(P * 100)
	diff := m - p
	for i, c := range coins {
		if c > diff {
			continue
		}

		v := diff % c
		a := int(diff / c)
		res[i] = a
		if v == 0 {
			break
		}

		diff -= int64(a) * c
	}

	res2 := make([]int, 6)
	c := 0
	for i := len(res) - 1; i >= 0; i-- {
		res2[c] = res[i]
		c++
	}

	return res2
}
