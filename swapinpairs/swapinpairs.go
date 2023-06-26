package swapinpairs

import "github.com/gonzispina/gorithms/helpers"

func swapHeadAndNext(head **helpers.ListNode) {
	n := (*head).Next
	(*head).Next = n.Next
	n.Next = *head
	*head = n
}

func swapPairs(head *helpers.ListNode) *helpers.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	swapHeadAndNext(&head)
	prev := head.Next
	n := head.Next.Next
	for n != nil && n.Next != nil {
		swapHeadAndNext(&n)
		prev.Next = n
		prev = prev.Next.Next
		n = n.Next.Next
	}

	return head
}
