package swapinpairs

import (
	"github.com/gonzispina/gorithms/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		res := swapPairs(helpers.NewLinkedList([]int{1}))
		assert.Equal(t, []int{1}, res.ToArray())
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := swapPairs(helpers.NewLinkedList([]int{1, 2}))
		assert.Equal(t, []int{2, 1}, res.ToArray())
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := swapPairs(helpers.NewLinkedList([]int{1, 2, 3, 4}))
		assert.Equal(t, []int{2, 1, 4, 3}, res.ToArray())
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := swapPairs(helpers.NewLinkedList([]int{1, 2, 3, 4, 5}))
		assert.Equal(t, []int{2, 1, 4, 3, 5}, res.ToArray())
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := swapPairs(helpers.NewLinkedList([]int{1, 2, 3, 4, 5, 6}))
		assert.Equal(t, []int{2, 1, 4, 3, 6, 5}, res.ToArray())
	})

	t.Run("Test Solution", func(t *testing.T) {
		res := swapPairs(helpers.NewLinkedList([]int{1, 2, 3, 4, 5, 6, 7}))
		assert.Equal(t, []int{2, 1, 4, 3, 6, 5, 7}, res.ToArray())
	})
}
