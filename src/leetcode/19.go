package leetcode

func Problem19() {

}

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	tail := head
	for i := 0; i < n; i++ {
		tail = tail.Next
	}
	if tail == nil {
		return head.Next
	}

	current := head
	for tail.Next != nil {
		tail = tail.Next
		current = current.Next
	}
	current.Next = current.Next.Next
	return head
}
