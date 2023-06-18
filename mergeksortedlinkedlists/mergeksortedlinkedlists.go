package insertionsortsinglell

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

func mergeKLists(lists []*helpers.ListNode) *helpers.ListNode {
	if len(lists) == 0 {
		return nil
	}

	for len(lists) > 1 {
		l1 := lists[0]
		l2 := lists[1]

		lists = lists[2:]
		merge2Lists(&l1, &l2)

		lists = append(lists, l1)
	}

	return lists[0]
}
