package helpers

type ListNode struct {
	Val  int
	Next *ListNode
}

func (l *ListNode) ToArray() []int {
	if l == nil {
		return []int{}
	}

	var res []int
	n := l
	for n != nil {
		res = append(res, n.Val)
		n = n.Next
	}
	return res
}

func NewLinkedList(arr []int) *ListNode {
	if len(arr) == 0 {
		return nil
	}

	res := &ListNode{
		Val:  arr[0],
		Next: nil,
	}

	n := res
	for i := 1; i < len(arr); i++ {
		n.Next = &ListNode{Val: arr[i]}
		n = n.Next
	}

	return res
}