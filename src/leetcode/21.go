package leetcode

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}

	if l2 == nil {
		return l1
	}

	var head *ListNode

	if l1.Val <= l2.Val {
		head = l1
		l1 = l1.Next

	} else {
		head = l2
		l2 = l2.Next
	}
	current := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			current.Next = l1
			current = current.Next
			l1 = l1.Next

		} else {
			current.Next = l2
			current = current.Next
			l2 = l2.Next
		}
	}

	if l1 == nil {
		current.Next = l2
	} else {
		current.Next = l1
	}
	return head
}
