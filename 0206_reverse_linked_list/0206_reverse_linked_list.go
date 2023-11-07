package leetcode

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

type LinkedList struct {
	Head *ListNode
	Tail *ListNode
}

func NewLinkedList() *LinkedList {
	return &LinkedList{Head: nil, Tail: nil}
}

func (list *LinkedList) Append(val int) {
	if list.Head == nil {
		list.Head = &ListNode{Val: val, Next: nil}
		list.Tail = list.Head
		return
	}
	list.Tail.Next = &ListNode{Val: val, Next: nil}
	list.Tail = list.Tail.Next
}

func (list *LinkedList) Print() {
	curr := list.Head
	for curr != nil {
		fmt.Print(curr.Val)
		curr = curr.Next
	}
	fmt.Println()
}

func reverseList(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	curr := head
	var pre *ListNode
	for curr != nil {
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre
}

func main() {
	list := NewLinkedList()
	list.Append(1)
	list.Append(2)
	list.Append(3)
	list.Append(4)
	list.Append(5)
	list.Print()
	list.Head = reverseList(list.Head)
	list.Print()
}
