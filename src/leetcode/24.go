package leetcode

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	current := head
	post := current.Next
	current.Next = post.Next
	post.Next = current
	head = post
	pre := current
	current = pre.Next
	if current == nil {
		return head
	}
	post = current.Next
	if post == nil {
		return head
	}

	for current != nil && post != nil {
		current.Next = post.Next
		pre.Next = post
		post.Next = current
		pre = current
		current = pre.Next
		if current == nil {
			return head
		}
		post = current.Next
		if post == nil {
			return head
		}
	}
	return head
}
