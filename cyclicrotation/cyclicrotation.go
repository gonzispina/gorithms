package cyclicrotation

// CyclicRotation moves every value in the array A, K times to the right
// Solution in O(N)
func CyclicRotation(a []int, k int) []int {
	if a == nil || len(a) == 0 {
		return []int{}
	}

	if k >= len(a) {
		k = k % len(a)
	}

	backwards := k > len(a) / 2
	if backwards {
		k = len(a) - k
	}

	for i := 0; i < k; i++ {
		if backwards {
			prev := a[0]
			for j := len(a) - 1; j >= 0; j-- {
				value := a[j]
				a[j] = prev
				prev = value
			}
		} else {
			prev := a[len(a) - 1]
			for index, value := range a {
				a[index] = prev
				prev = value
			}
		}
	}

	return a
}

