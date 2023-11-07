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

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}
	newHead := &ListNode{}
	curr := newHead
	for list1 != nil && list2 != nil {
		if list1.Val < list2.Val {
			curr.Next = list1
			list1 = list1.Next
		} else {
			curr.Next = list2
			list2 = list2.Next
		}
		curr = curr.Next
	}
	if list1 == nil {
		curr.Next = list2
	}
	if list2 == nil {
		curr.Next = list1
	}
	return newHead.Next
}

func main() {
	list1 := NewLinkedList()
	list1.Append(1)
	list1.Append(2)
	list1.Append(4)
	list2 := NewLinkedList()
	list2.Append(1)
	list2.Append(3)
	list2.Append(4)
	res := NewLinkedList()
	res.Head = mergeTwoLists(list1.Head, list2.Head)
	res.Print() // 112344
}
