package sortlinkedlist

import "github.com/gonzispina/gorithms/helpers"

func merge2Lists(l1, l2 **helpers.ListNode) {
	if *l2 == nil {
		return
	}

	if *l1 == nil || (*l2).Val < (*l1).Val {
		*l1, *l2 = *l2, *l1
	}

	n := *l1
	m := *l2
	for m != nil && n.Next != nil {
		if n.Next.Val < m.Val {
			n = n.Next
			continue
		}

		tmp := m.Next
		m.Next = n.Next
		n.Next = m
		m = tmp
	}

	if m != nil {
		n.Next = m
	}
}

func sortList(head *helpers.ListNode) *helpers.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	if head.Next.Next == nil {
		if head.Val > head.Next.Val {
			head.Next.Next = head
			head = head.Next
			head.Next.Next = nil
			return head
		}

		return head
	}

	fast := head
	slow := head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil

	head = sortList(head)
	tmp = sortList(tmp)

	merge2Lists(&head, &tmp)

	return head
}
