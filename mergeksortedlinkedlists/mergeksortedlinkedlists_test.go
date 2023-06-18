package insertionsortsinglell

import (
	"github.com/gonzispina/gorithms/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution 1", func(t *testing.T) {
		lists := []*helpers.ListNode{
			helpers.NewLinkedList([]int{-1, 0, 1, 2, 3, 4}),
			helpers.NewLinkedList([]int{2, 3, 4, 5, 6, 7, 8}),
			helpers.NewLinkedList([]int{-3, -2, -1}),
		}

		res := mergeKLists(lists)
		assert.Equal(t, []int{-3, -2, -1, -1, 0, 1, 2, 2, 3, 3, 4, 4, 5, 6, 7, 8}, res.ToArray())
	})

	t.Run("Test Solution 2", func(t *testing.T) {
		lists := []*helpers.ListNode{
			helpers.NewLinkedList([]int{1, 4, 5}),
			helpers.NewLinkedList([]int{1, 3, 4}),
			helpers.NewLinkedList([]int{2, 6}),
		}

		res := mergeKLists(lists)
		assert.Equal(t, []int{1, 1, 2, 3, 4, 4, 5, 6}, res.ToArray())
	})

	t.Run("Test Solution 3", func(t *testing.T) {
		lists := []*helpers.ListNode{
			helpers.NewLinkedList([]int{1, 4, 5}),
			helpers.NewLinkedList([]int{1, 3, 4}),
			helpers.NewLinkedList([]int{-2, 6}),
			helpers.NewLinkedList([]int{0}),
			helpers.NewLinkedList([]int{-6, -2, 1, 4}),
		}

		res := mergeKLists(lists)
		assert.Equal(t, []int{-6, -2, -2, 0, 1, 1, 1, 3, 4, 4, 4, 5, 6}, res.ToArray())
	})

	t.Run("Test Solution 4", func(t *testing.T) {
		lists := []*helpers.ListNode{
			helpers.NewLinkedList([]int{1, 4, 5}),
			helpers.NewLinkedList([]int{1, 3, 4}),
			helpers.NewLinkedList([]int{-2, 6}),
			helpers.NewLinkedList([]int{0, 0}),
			helpers.NewLinkedList([]int{-6, -2, 1, 4}),
			helpers.NewLinkedList([]int{-10, -8}),
		}

		res := mergeKLists(lists)
		assert.Equal(t, []int{-10, -8, -6, -2, -2, 0, 0, 1, 1, 1, 3, 4, 4, 4, 5, 6}, res.ToArray())
	})

	t.Run("Test Solution 4", func(t *testing.T) {
		lists := []*helpers.ListNode{
			helpers.NewLinkedList([]int{0}),
			helpers.NewLinkedList([]int{1}),
			helpers.NewLinkedList([]int{3}),
			helpers.NewLinkedList([]int{6}),
		}

		res := mergeKLists(lists)
		assert.Equal(t, []int{0, 1, 3, 6}, res.ToArray())
	})

	t.Run("Test Solution empty lists", func(t *testing.T) {
		lists := []*helpers.ListNode{
			helpers.NewLinkedList([]int{}),
			helpers.NewLinkedList([]int{}),
		}

		res := mergeKLists(lists)
		assert.Equal(t, []int{}, res.ToArray())
	})

	t.Run("Test Solution empty lists", func(t *testing.T) {
		lists := []*helpers.ListNode{
			helpers.NewLinkedList([]int{-1}),
			helpers.NewLinkedList([]int{}),
			helpers.NewLinkedList([]int{1, 2}),
		}

		res := mergeKLists(lists)
		assert.Equal(t, []int{-1, 1, 2}, res.ToArray())
	})
}
