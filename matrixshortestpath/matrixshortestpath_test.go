package matrixshortestpath

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution N= ", func(t *testing.T) {
		res := shortestPathBinaryMatrix([][]int{{0, 0, 0}, {1, 1, 0}, {1, 1, 0}})
		assert.Equal(t, 4, res)
	})
}
