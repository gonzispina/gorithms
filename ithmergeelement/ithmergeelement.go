package ithmergeelement

func findIthMergeElement(a, b []int, i int) int {
	if a[0] > b[0] {
		a, b = b, a
	}

	aLeft := 0
	aRight := len(a) - 1
	bLeft := 0
	//bRight := len(b) - 1
	for i-aLeft-bLeft > 0 {
		for aLeft < aRight {

		}
	}
	return 0
}

// 1, 3, 5, 7, 9         4, 6, 8, 10           5
// log(A)
