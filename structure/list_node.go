package structrue

type ListNode struct {
	Val  int
	Next *ListNode
}

// ListToInts convert ListNode to []int
func ListToInts(head *ListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// IntsToList convert []int to ListNode
func IntsToList(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}

	head := &ListNode{}
	curr := head
	for _, num := range nums {
		curr.Next = &ListNode{Val: num}
		curr = curr.Next
	}

	return head.Next
}
