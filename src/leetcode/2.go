package leetcode

import "fmt"

func Problem2() {
	//node3 := ListNode{Val:3, Next:nil}
	node4 := ListNode{Val: 8, Next: nil}
	node2 := ListNode{Val: 1, Next: &node4}

	//n4 := ListNode{Val:4, Next:nil}
	//n6 := ListNode{Val:6, Next:&n4}
	n5 := ListNode{Val: 0, Next: nil}

	showList(&node2)
	showList(&n5)
	res := addTwoNumbers(&node2, &n5)
	showList(res)
}

func showList(head *ListNode) {

	if head == nil {
		fmt.Printf("NULL\n")
		return
	}
	if head.Next == nil {
		fmt.Printf("%d\n", head.Val)
		return
	}

	current := head
	for current.Next != nil {
		fmt.Printf("%d --> ", current.Val)
		current = current.Next
	}
	if current.Next == nil {
		fmt.Printf("%d\n", current.Val)
	}

	return
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	if l1 == nil && l2 == nil {
		return nil
	}

	if l1 == nil && l2 != nil {
		return l2
	}

	if l2 == nil && l1 != nil {
		return l1
	}

	var head *ListNode
	var current = head
	var flag int
	for {
		if l1 == nil || l2 == nil {
			break
		}

		val := l1.Val + l2.Val + flag
		if val >= 10 {
			flag = 1
		} else {
			flag = 0
		}
		val = val % 10
		node := ListNode{
			Val:  val,
			Next: nil,
		}
		if head != nil {
			current.Next = &node
			current = current.Next
		} else {
			head = &node
			current = head
		}

		l1 = l1.Next
		l2 = l2.Next

	}

	if l1 != nil && l2 == nil {
		for l1 != nil {
			//fmt.Println("!!!!!")
			val := l1.Val + flag
			if val >= 10 {
				flag = 1
			} else {
				flag = 0
			}
			val = val % 10
			node := ListNode{
				Val:  val,
				Next: nil,
			}
			//fmt.Println(node)
			current.Next = &node
			current = current.Next
			l1 = l1.Next
		}
	}
	if l2 != nil && l1 == nil {
		for l2 != nil {
			val := l2.Val + flag
			if val >= 10 {
				flag = 1
			} else {
				flag = 0
			}
			val = val % 10
			node := ListNode{
				Val:  val,
				Next: nil,
			}
			current.Next = &node
			current = current.Next
			l2 = l2.Next
		}
	}
	if flag == 1 {
		node := ListNode{
			Val:  1,
			Next: nil,
		}
		current.Next = &node
		current = current.Next
	}
	return head
}
