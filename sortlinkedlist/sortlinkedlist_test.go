package sortlinkedlist

import (
	"github.com/gonzispina/gorithms/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test sorting list", func(t *testing.T) {
		res := sortList(helpers.NewLinkedList([]int{2, 3, 0, 0, -1, 4, 3, 2, 3, 4}))
		assert.Equal(t, []int{-1, 0, 0, 2, 2, 3, 3, 3, 4, 4}, res.ToArray())
	})
}
