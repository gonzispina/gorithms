package moretotheleft

func moreToTheLeft(arr []int) bool {
	q := [][]int{arr}

	for len(q[0]) > 2 {
		a := q[0]
		q = q[1:]

		q = append(q, a[0:len(a)/2])
		q = append(q, a[len(a)/2:])
	}

	for len(q) > 1 {
		a1 := q[0]
		a2 := q[1]
		q = q[2:]

		if a1[0] < a1[1] || a2[0] < a2[1] {
			return false
		}

		q = append(q, []int{a1[0] + a1[1], a2[0] + a2[1]})
	}

	return q[0][0] > q[0][1]
}
