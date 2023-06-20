package binarytreemaxnodedistance

import (
	"github.com/gonzispina/gorithms/helpers"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestSolution(t *testing.T) {
	t.Run("Test Solution", func(t *testing.T) {
		res := distance(&helpers.TreeNode{
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
		})

		assert.Equal(t, 2, res)
	})
} 
