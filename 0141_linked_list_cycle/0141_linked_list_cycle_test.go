package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func createCycleList(nums []int, pos int) *ListNode {
	head := structure.IntsToList(nums)
	if pos == -1 {
		return head
	}

	curr := head
	for pos > 0 {
		curr = curr.Next
		pos--
	}

	tail := head
	var preTail *ListNode
	for tail != nil {
		preTail = tail
		tail = tail.Next
	}
	preTail.Next = curr

	return head
}

func TestHasCycle(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		head := createCycleList([]int{3, 2, 0, -4}, 1)
		assert.True(t, hasCycle(head))
	})

	t.Run("2", func(t *testing.T) {
		head := createCycleList([]int{1, 2}, 0)
		assert.True(t, hasCycle(head))
	})

	t.Run("3", func(t *testing.T) {
		head := createCycleList([]int{1}, -1)
		assert.False(t, hasCycle(head))
	})

	t.Run("4", func(t *testing.T) {
		head := createCycleList([]int{}, -1)
		assert.False(t, hasCycle(head))
	})
}
