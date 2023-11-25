package leetcode

import (
	"github.com/leonasdev/leetcode/structure"
)

type ListNode = structure.ListNode

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{Next: head}
	l, r := dummyNode, head

	for i := 0; i < n; i++ {
		r = r.Next
	}

	for r != nil {
		r = r.Next
		l = l.Next
	}

	l.Next = l.Next.Next

	return dummyNode.Next
}
