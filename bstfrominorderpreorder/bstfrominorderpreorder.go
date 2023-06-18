package bstfrominorderpreorder

import "github.com/gonzispina/gorithms/helpers"

func intToLeaf(i int) *helpers.TreeNode {
	return &helpers.TreeNode{
		Val: i,
	}
}

func buildSubTree(preorder []int, inorder []int) *helpers.TreeNode {
	if len(inorder) == 0 {
		return nil
	}

	if len(inorder) == 1 {
		return intToLeaf(preorder[0])
	}

	root := intToLeaf(preorder[0])
	if preorder[0] == inorder[0] {
		root.Right = intToLeaf(preorder[1])
	} else {
		root.Left = intToLeaf(preorder[1])
	}
	return root
}

func buildTree(preorder []int, inorder []int) *helpers.TreeNode {
	if len(preorder) <= 2 {
		return buildSubTree(preorder, inorder)
	}

	value := preorder[0]
	preorder = preorder[1:]

	var i int
	for i = 0; i < len(inorder); i++ {
		if inorder[i] == value {
			break
		}
	}

	root := intToLeaf(value)
	root.Left = buildTree(preorder[:i], inorder[:i])
	root.Right = buildTree(preorder[i:], inorder[i+1:])

	return root
}
