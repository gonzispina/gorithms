package findfalseinmatrix

// Excercise says to assume this op as atomic
func hasFalse(m [][]bool, i0, j0, i1, j1 int) bool {
	for i := i0; i < i1; i++ {
		for j := j0; j < j1; j++ {
			if m[i][j] == false {
				return true
			}
		}
	}

	return false
}

func findFalseInMatrixAux(m [][]bool, i0, j0, i1, j1 int) (int, int) {
	if i1-i0 <= 1 && j1-j0 <= 1 {
		if m[i0][j0] {
			return -1, -1
		}

		return i0, j0
	}

	if !hasFalse(m, i0, j0, i1, j1) {
		return -1, -1
	}

	if i, j := findFalseInMatrixAux(m, i0, j0, (i1)/2, (j1)/2); i >= 0 {
		return i, j
	} else if i, j := findFalseInMatrixAux(m, (i0+i1)/2, j0, i1, (j1)/2); i >= 0 {
		return i, j
	} else if i, j := findFalseInMatrixAux(m, i0, (j0+j1)/2, (i1)/2, j1); i >= 0 {
		return i, j
	} else {
		return findFalseInMatrixAux(m, (i0+i1)/2, (j0+j1)/2, i1, j1)
	}
}

func findFalseInMatrix(m [][]bool) (int, int) {
	return findFalseInMatrixAux(m, 0, 0, len(m), len(m))
}
