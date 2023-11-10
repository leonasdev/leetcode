package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func newListNode(val int) *ListNode {
	return &ListNode{
		Val:  val,
		Next: nil,
	}
}

func TestReverseList(t *testing.T) {
	root := &ListNode{
		Val: 1,
	}
	curr := root
	for i := 2; i <= 5; i++ {
		curr.Next = newListNode(i)
		curr = curr.Next
	}

	root = reverseList(root)

	want := []int{5, 4, 3, 2, 1}

	curr = root
	for _, w := range want {
		assert.Equal(t, w, curr.Val)
		curr = curr.Next
	}
}
