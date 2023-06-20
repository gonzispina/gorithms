package countunorderedpairs

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution A = 3 2 1", func(t *testing.T) {
		arr := []int{3, 2, 1}
		_, res := count(arr)
		assert.Equal(t, 3, res)
	})

	t.Run("Test Solution A = 4 3 2 1", func(t *testing.T) {
		arr := []int{4, 3, 2, 1}
		_, res := count(arr)
		assert.Equal(t, 6, res)
	})

	t.Run("Test Solution A = 5 4 3 2 1", func(t *testing.T) {
		arr := []int{5, 4, 3, 2, 1}
		_, res := count(arr)
		assert.Equal(t, 10, res)
	})

	t.Run("Test Solution A = 6 5 4 3 2 1", func(t *testing.T) {
		arr := []int{6, 5, 4, 3, 2, 1}
		_, res := count(arr)
		assert.Equal(t, 15, res)
	})

	t.Run("Test Solution A = 9 3 3 5 6 0 1 4", func(t *testing.T) {
		arr := []int{9, 0, 1, 3, 3, 5, 6, 4}
		_, res := count(arr)
		assert.Equal(t, 9, res)
	})

	t.Run("Test Solution A = 3 2 1 4", func(t *testing.T) {
		arr := []int{3, 2, 1, 4}
		_, res := count(arr)
		assert.Equal(t, 3, res)
	})

	t.Run("Test Solution A = 3 2 5 1 4", func(t *testing.T) {
		arr := []int{3, 2, 5, 1, 4}
		_, res := count(arr)
		assert.Equal(t, 5, res)
	})

	t.Run("Test Solution A = 9 3 3 5 6 0 1 4 4 4", func(t *testing.T) {
		arr := []int{9, 0, 1, 3, 3, 5, 6, 4, 4, 4}
		_, res := count(arr)
		assert.Equal(t, 15, res)
	})

	t.Run("Test Solution A = 0 0 0 0", func(t *testing.T) {
		arr := []int{0, 0, 0, 0}
		_, res := count(arr)
		assert.Equal(t, 0, res)
	})

	t.Run("Test Solution A = 0 1 2 3", func(t *testing.T) {
		arr := []int{0, 1, 2, 3}
		_, res := count(arr)
		assert.Equal(t, 0, res)
	})
}
