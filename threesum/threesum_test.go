package threesum

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution 0, 0, 0", func(t *testing.T) {
		res := threeSum([]int{0, 0, 0})
		assert.Equal(t, [][]int{{0, 0, 0}}, res)

	})

	t.Run("Test Solution -4, -1, -1, 0, 1, 2", func(t *testing.T) {
		res := threeSum([]int{-1, 0, 1, 2, -1, -4})
		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, res)
	})

	t.Run("Test Solution -4, -1, -1, 0, 1, 2, 4", func(t *testing.T) {
		res := threeSum([]int{-1, 0, 1, 2, -1, -4, 4})
		assert.Equal(t, [][]int{{-4, 0, 4}, {-1, -1, 2}, {-1, 0, 1}}, res)
	})

	t.Run("Test Solution -4, -1, -1, 0, 1, 1, 2", func(t *testing.T) {
		res := threeSum([]int{-1, 0, 1, 2, -1, -4, 1})
		assert.Equal(t, [][]int{{-1, -1, 2}, {-1, 0, 1}}, res)
	})

	t.Run("Test Solution -4, -1, -1, 0, 1, 1, 2, 4", func(t *testing.T) {
		res := threeSum([]int{-1, 0, 1, 2, -1, -4, 1, 4})
		assert.Equal(t, [][]int{{-4, 0, 4}, {-1, -1, 2}, {-1, 0, 1}}, res)
	})

	t.Run("Test Solution -2, 0, 1, 1, 2", func(t *testing.T) {
		res := threeSum([]int{-2, 0, 1, 1, 2})
		assert.Equal(t, [][]int{{-2, 0, 2}, {-2, 1, 1}}, res)
	})

	t.Run("Test Solution -4, -3, -2, -1, -1, 0, 0, 1, 2, 3, 4", func(t *testing.T) {
		res := threeSum([]int{-1, 0, 1, 2, -1, -4, -2, -3, 3, 0, 4})
		assert.Equal(t, [][]int{{-4, 0, 4}, {-4, 1, 3}, {-3, -1, 4}, {-3, 0, 3}, {-3, 1, 2}, {-2, -1, 3}, {-2, 0, 2}, {-1, -1, 2}, {-1, 0, 1}}, res)
	})

}
