package findfalseinmatrix

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution 2x2 i = 0, j = 0", func(t *testing.T) {
		m := [][]bool{
			{false, true},
			{true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 0, i)
		assert.Equal(t, 0, j)
	})

	t.Run("Test Solution 2x2 i = 0, j = 1", func(t *testing.T) {
		m := [][]bool{
			{true, false},
			{true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
	})

	t.Run("Test Solution 3x3 i = 0, j = i", func(t *testing.T) {
		m := [][]bool{
			{true, false, true},
			{true, true, true},
			{true, true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 0, i)
		assert.Equal(t, 1, j)
	})

	t.Run("Test Solution 3x3 i=1, j=1", func(t *testing.T) {
		m := [][]bool{
			{true, true, true},
			{true, false, true},
			{true, true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 1, i)
		assert.Equal(t, 1, j)
	})

	t.Run("Test Solution 3x3 i=1, j=0", func(t *testing.T) {
		m := [][]bool{
			{true, true, true},
			{false, true, true},
			{true, true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 1, i)
		assert.Equal(t, 0, j)
	})

	t.Run("Test Solution 4x4 i = 2, j = 2", func(t *testing.T) {
		m := [][]bool{
			{true, true, true, true},
			{true, true, true, true},
			{true, true, false, true},
			{true, true, true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 2, i)
		assert.Equal(t, 2, j)
	})

	t.Run("Test Solution 4x4 i = 0, j = 0", func(t *testing.T) {
		m := [][]bool{
			{false, true, true, true},
			{true, true, true, true},
			{true, true, true, true},
			{true, true, true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 0, i)
		assert.Equal(t, 0, j)
	})

	t.Run("Test Solution 4x4 i=3, j=3", func(t *testing.T) {
		m := [][]bool{
			{true, true, true, true},
			{true, true, true, true},
			{true, true, true, true},
			{true, true, true, false},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 3, i)
		assert.Equal(t, 3, j)
	})

	t.Run("Test Solution 4x4 i=1, j=3", func(t *testing.T) {
		m := [][]bool{
			{true, true, true, true},
			{true, true, true, false},
			{true, true, true, true},
			{true, true, true, true},
		}

		i, j := findFalseInMatrix(m)
		assert.Equal(t, 1, i)
		assert.Equal(t, 3, j)
	})
}
