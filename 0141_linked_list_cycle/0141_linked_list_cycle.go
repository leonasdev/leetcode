package leetcode

import "github.com/leonasdev/leetcode/structure"

type ListNode = structure.ListNode

// O(1) memory
func hasCycle(head *ListNode) bool {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			return true
		}
	}
	return false
}

// O(n) memory
func hasCycle2(head *ListNode) bool {
	m := map[*ListNode]struct{}{}

	for head != nil {
		if _, exists := m[head]; exists {
			return true
		}
		m[head] = struct{}{}
		head = head.Next
	}

	return false
}
