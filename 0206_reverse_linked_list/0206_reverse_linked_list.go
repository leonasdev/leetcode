package leetcode

import structrue "github.com/leonasdev/leetcode/structure"

type ListNode = structrue.ListNode

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
