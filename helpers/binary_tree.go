package helpers

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

//te amoooo

func (n *TreeNode) InorderTree2Arr(res *[]int) {
	if n == nil {
		return
	}

	n.Left.InorderTree2Arr(res)
	*res = append(*res, n.Val)
	n.Right.InorderTree2Arr(res)
}

func (n *TreeNode) PreorderTree2Arr(res *[]int) {
	if n == nil {
		return
	}

	*res = append(*res, n.Val)
	n.Left.InorderTree2Arr(res)
	n.Right.InorderTree2Arr(res)
}

func (n *TreeNode) PostorderTree2Arr(res *[]int) {
	if n == nil {
		return
	}

	n.Left.InorderTree2Arr(res)
	n.Right.InorderTree2Arr(res)
	*res = append(*res, n.Val)
}
