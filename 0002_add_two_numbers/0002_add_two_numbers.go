package leetcode

import "github.com/leonasdev/leetcode/structure"

type ListNode = structure.ListNode

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	curr := dummy
	v1, v2, carry := 0, 0, 0

	for l1 != nil || l2 != nil || carry != 0 {
		if l1 != nil {
			v1 = l1.Val
			l1 = l1.Next
		} else {
			v1 = 0
		}
		if l2 != nil {
			v2 = l2.Val
			l2 = l2.Next
		} else {
			v2 = 0
		}

		sum := (v1 + v2 + carry) % 10
		carry = (v1 + v2 + carry) / 10
		curr.Next = &ListNode{Val: sum}
		curr = curr.Next
	}

	return dummy.Next
}
