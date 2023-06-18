package sortedarr2bst

import "github.com/gonzispina/gorithms/helpers"

func intToLeaf(i int) *helpers.TreeNode {
	return &helpers.TreeNode{
		Val: i,
	}
}

func buildSubTree(nums []int) *helpers.TreeNode {
	if len(nums) == 0 {
		return nil
	}

	if len(nums) == 1 {
		return intToLeaf(nums[0])
	}

	root := intToLeaf(nums[1])
	root.Left = intToLeaf(nums[0])

	if len(nums) == 2 {
		return root
	}

	root.Right = intToLeaf(nums[2])
	return root
}

func sortedArrayToBST(nums []int) *helpers.TreeNode {
	if len(nums) <= 3 {
		return buildSubTree(nums)
	}

	mid := len(nums) / 2
	root := intToLeaf(nums[mid])
	root.Left = sortedArrayToBST(nums[0:mid])
	root.Right = sortedArrayToBST(nums[mid+1:])
	return root
}
