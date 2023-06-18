package disjointunionofintervals

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test ordered", func(t *testing.T) {
		res := merge([][]int{{1, 6}, {7, 8}, {8, 11}, {9, 10}})
		assert.Equal(t, [][]int{{1, 6}, {7, 11}}, res)
	})

	t.Run("Test unordered", func(t *testing.T) {
		res := merge([][]int{{9, 10}, {7, 8}, {1, 6}, {8, 11}})
		assert.Equal(t, [][]int{{1, 6}, {7, 11}}, res)
	})

	t.Run("Test 1", func(t *testing.T) {
		res := merge([][]int{{9, 10}})
		assert.Equal(t, [][]int{{9, 10}}, res)
	})

	t.Run("Test 2", func(t *testing.T) {
		res := merge([][]int{{9, 10}, {16, 17}})
		assert.Equal(t, [][]int{{9, 10}, {16, 17}}, res)
	})

	t.Run("Test overlapping", func(t *testing.T) {
		res := merge([][]int{{9, 20}, {16, 33}})
		assert.Equal(t, [][]int{{9, 33}}, res)
	})

	t.Run("Test with negatives", func(t *testing.T) {
		res := merge([][]int{{-1, 20}, {16, 33}})
		assert.Equal(t, [][]int{{-1, 33}}, res)
	})

	t.Run("Test single points", func(t *testing.T) {
		res := merge([][]int{{-1, 20}, {0, 0}, {16, 33}})
		assert.Equal(t, [][]int{{-1, 33}}, res)
	})

	t.Run("Test single points 2", func(t *testing.T) {
		res := merge([][]int{{-1, 0}, {0, 0}, {1, 1}})
		assert.Equal(t, [][]int{{-1, 0}, {1, 1}}, res)
	})

	t.Run("Test single points 2", func(t *testing.T) {
		res := merge([][]int{{-2, 0}, {0, 0}, {-1, 100}})
		assert.Equal(t, [][]int{{-2, 100}}, res)
	})
}
