package leetcode

import (
	"github.com/leonasdev/leetcode/structure"
)

type ListNode = structure.ListNode

func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode
	for cur != nil {
		tmp := cur.Next
		cur.Next = pre
		pre = cur
		cur = tmp
	}
	return pre
}
