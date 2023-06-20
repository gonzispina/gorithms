package fixedpointsortedarray

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution −4, −1, 2, 4, 7", func(t *testing.T) {
		res := fixedPoint([]int{-4, -1, 2, 4, 7})
		assert.Equal(t, 4, res)
	})

	t.Run("Test Solution −4, −1, 2, 5, 5, 7", func(t *testing.T) {
		res := fixedPoint([]int{-4, -1, 2, 5, 5, 7})
		assert.Equal(t, 5, res)
	})

	t.Run("Test Solution 0, 2, 10, 11, 12, 13", func(t *testing.T) {
		res := fixedPoint([]int{0, 2, 10, 11, 12, 13})
		assert.Equal(t, 2, res)
	})

	t.Run("Test Solution 1, 3, 10, 11, 12, 13", func(t *testing.T) {
		res := fixedPoint([]int{1, 3, 10, 11, 12, 13})
		assert.Equal(t, 1, res)
	})

	t.Run("Test Solution 0, 3, 4, 5, 6, 7", func(t *testing.T) {
		res := fixedPoint([]int{0, 3, 4, 5, 6, 7})
		assert.Equal(t, -1, res)
	})

	t.Run("Test Solution 0, 3, 4, 5, 6, 7, 7, 7", func(t *testing.T) {
		res := fixedPoint([]int{-1, -2, 2, 3, 4, 5, 6, 7, 9})
		assert.Equal(t, 9, res)
	})
}
