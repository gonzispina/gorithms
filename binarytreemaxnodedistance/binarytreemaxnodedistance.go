package binarytreemaxnodedistance

import "github.com/gonzispina/gorithms/helpers"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func distance(n *helpers.TreeNode) int {
	if n == nil || (n.Right == nil && n.Left == nil) {
		return 0
	}

	return max(distance(n.Left), distance(n.Right)) + 1
}
