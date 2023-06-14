package matrixshortestpath

func calculateCost(grid [][]int, visited [][]bool, i, j int) int {
	if grid[i][j] == 1 {
		return -1
	}

	if i == len(grid)-1 && j == len(grid)-1 {
		return 1
	}

	visited[i][j] = true
	minCost := -1
	for i0 := i - 1; i0 < i+2; i0++ {
		for j0 := j - 1; j0 < j+2; j0++ {
			if !validPosition(grid, visited, i0, j0) {
				continue
			}

			cost := calculateCost(grid, visited, i0, j0)
			if cost > -1 {
				if cost == 1 {
					visited[i][j] = false
					return 2
				}
				if minCost == -1 || cost < minCost {
					minCost = cost
				}
			}
		}
	}

	visited[i][j] = false

	if minCost == -1 {
		return -1
	}

	return minCost + 1
}

func shortestPathBinaryMatrixRecursive(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid))
	}

	return calculateCost(grid, visited, 0, 0)
}

func validPosition(grid [][]int, visited [][]bool, i, j int) bool {
	return i >= 0 && j >= 0 && i < len(grid) && j < len(grid) && grid[i][j] == 0 && !visited[i][j]
}

func shortestPathBinaryMatrix(grid [][]int) int {
	visited := make([][]bool, len(grid))
	for i := 0; i < len(grid); i++ {
		visited[i] = make([]bool, len(grid))
	}

	if grid[0][0] == 1 {
		return -1
	}

	queue := [][]int{{0, 0, 0}}
	for len(queue) > 0 {
		i, j, cost := queue[0][0], queue[0][1], queue[0][2]
		queue = queue[1:]

		visited[i][j] = true
		cost++

		if i == len(grid)-1 && j == len(grid)-1 {
			// The first one to reach the goal will be the one that traveled
			// the less distance
			return cost
		}

		for i0 := i - 1; i0 < i+2; i0++ {
			for j0 := j - 1; j0 < j+2; j0++ {
				if !validPosition(grid, visited, i0, j0) {
					continue
				}

				queue = append(queue, []int{i0, j0, cost})
			}
		}
	}

	return -1
}
