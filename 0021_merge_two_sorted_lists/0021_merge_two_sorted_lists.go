package leetcode

import "github.com/leonasdev/leetcode/structure"

type ListNode = structure.ListNode

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	mergeList := &ListNode{}
	curr := mergeList
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
	return mergeList.Next
}
