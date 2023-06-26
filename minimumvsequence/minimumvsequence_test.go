package minimumvsequence

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		seq := []int{4, 2, -3, 5}
		res := findMinimum(seq)
		assert.Equal(t, -3, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		seq := []int{4, 2, -3, 5}
		res := findMinimum(seq)
		assert.Equal(t, -3, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		seq := []int{4, 2, -3, 5, 6}
		res := findMinimum(seq)
		assert.Equal(t, -3, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		seq := []int{-3, 5, 6, 7, 10, 20, 30}
		res := findMinimum(seq)
		assert.Equal(t, -3, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		seq := []int{4, 1, 5, 6, 7, 10, 20, 30}
		res := findMinimum(seq)
		assert.Equal(t, 1, res)
	})

	t.Run("Test Solution", func(t *testing.T) {
		seq := []int{20, 10, 6, 5, 4, 3, 1}
		res := findMinimum(seq)
		assert.Equal(t, 1, res)
	})
}
