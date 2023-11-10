package leetcode

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

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

func (list *LinkedList) toString() string {
	res := ""
	curr := list.Head
	for curr != nil {
		res += fmt.Sprint(curr.Val)
		curr = curr.Next
	}
	return res
}

func TestMergeTwoLists(t *testing.T) {
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
	assert.Equal(t, "112344", res.toString())
}
