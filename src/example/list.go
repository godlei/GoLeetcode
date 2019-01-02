package example

import "fmt"

type LinkNode struct {
	Val  int
	Next *LinkNode
}

func ListExample() {
	head := NewLinkList()

	addNode(1, 1, head)
	addNode(2, 2, head)
	addNode(3, 3, head)
	showLinkedList(head)
}

func NewLinkList() (head *LinkNode) {
	head = new(LinkNode)
	return
}

func addNode(val int, pos int, head *LinkNode) {
	if head == nil {
		return
	}
	var i int
	current := head
	for i = 1; i < pos && current != nil; i++ {
		current = current.Next
	}

	if i < pos {
		fmt.Println("长度不够")
		return
	}

	nNode := LinkNode{
		Val: val,
	}
	nNode.Next = current.Next
	current.Next = &nNode
	return
}

func showLinkedList(head *LinkNode) {
	if head == nil {
		fmt.Println("linkedList not exists")
		return
	}
	if head.Next == nil {
		fmt.Println("NULL")
		return
	}
	current := head.Next
	for {
		if current.Next != nil {
			fmt.Printf("%d --> ", current.Val)
			current = current.Next
		} else {
			fmt.Printf("%d\n", current.Val)
			break
		}
	}
}
