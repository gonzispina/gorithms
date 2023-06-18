package insertionsortsinglell

type ListNode struct {
	Val  int
	Next *ListNode
}

func insertSorted(head **ListNode, n *ListNode) {
	if (*head).Val > n.Val {
		n.Next = *head
		*head = n
		return
	}

	m := *head
	for m.Next != nil {
		if m.Next.Val > n.Val {
			break
		}

		m = m.Next
	}

	n.Next = m.Next
	m.Next = n
}

func insertionSortList(head *ListNode) *ListNode {
	n := head
	for n != nil && n.Next != nil {
		if n.Next.Val >= n.Val {
			n = n.Next
			continue
		}

		m := n.Next
		n.Next = m.Next
		insertSorted(&head, m)
	}

	return head
}
