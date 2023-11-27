package leetcode

import (
	"github.com/leonasdev/leetcode/structure"
)

type ListNode = structure.ListNode

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	for head != nil {
		tmp := head.Next
		head.Next = prev
		prev = head
		head = tmp
	}

	return prev
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	tmp := slow.Next
	slow.Next = nil
	second := reverseList(tmp)
	first := head

	for second != nil {
		tmp = first.Next
		first.Next = second
		tmp2 := second.Next
		second.Next = tmp
		first = tmp
		second = tmp2
	}
}
