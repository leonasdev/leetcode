package leetcode

import (
	"github.com/leonasdev/leetcode/structure"
)

type ListNode = structure.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{Next: head}
	slow, fast := dummy, head
	for ; n > 0; n-- {
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	slow.Next = slow.Next.Next

	return dummy.Next
}
