package bstfrominorderpreorder

import (
	"github.com/gonzispina/gorithms/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test tree", func(t *testing.T) {
		tree := buildTree([]int{3, 9, 20, 15, 7}, []int{9, 3, 15, 20, 7})

		expected := &helpers.TreeNode{
			Val: 3,
			Left: &helpers.TreeNode{
				Val: 9,
			},
			Right: &helpers.TreeNode{
				Val: 20,
				Left: &helpers.TreeNode{
					Val: 15,
				},
				Right: &helpers.TreeNode{
					Val: 7,
				},
			},
		}
		assert.Equal(t, expected, tree)
	})

	t.Run("Test tree", func(t *testing.T) {
		tree := buildTree([]int{1, 2, 3}, []int{3, 2, 1})

		expected := &helpers.TreeNode{
			Val: 1,
			Left: &helpers.TreeNode{
				Val: 2,
				Left: &helpers.TreeNode{
					Val: 3,
				},
			},
		}
		assert.Equal(t, expected, tree)
	})
}
