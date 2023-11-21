package leetcode

import (
	structrue "github.com/leonasdev/leetcode/structure"
)

type ListNode = structrue.ListNode

func reverseList(head *ListNode) *ListNode {
	var pre *ListNode
	curr := head
	for curr != nil {
		tmp := curr.Next
		curr.Next = pre
		pre = curr
		curr = tmp
	}
	return pre
}

func reorderList(head *ListNode) {
	if head == nil {
		return
	}

	// find middle
	slow, fast := head, head.Next
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	second := slow.Next
	slow.Next = nil // first head's last should point to nil

	reverseHead := reverseList(second)

	// merge two halfs
	curr, rCurr := head, reverseHead
	for rCurr != nil {
		tmp := curr.Next
		curr.Next = rCurr
		rTmp := rCurr.Next
		rCurr.Next = tmp
		curr = tmp
		rCurr = rTmp
	}
}
