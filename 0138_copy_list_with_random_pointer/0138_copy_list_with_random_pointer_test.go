package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestCopyRandomList(t *testing.T) {
	nums := [][2]int{{7, -1}, {13, 0}, {11, 4}, {10, 2}, {1, 0}}
	head := structure.IntsToRandomList(nums)
	copyHead := copyRandomList(head)
	nodes := []*Node{}

	curr := head
	curr2 := copyHead
	for curr != nil {
		assert.Equal(t, curr.Val, curr2.Val)
		assert.NotSame(t, curr, curr2)
		if curr.Random != nil {
			assert.NotSame(t, curr.Random, curr2.Random)
		}

		nodes = append(nodes, curr2)
		curr = curr.Next
		curr2 = curr2.Next
	}
	assert.Nil(t, curr)

	curr2 = copyHead
	i := 0
	for curr2 != nil {
		index := nums[i][1]
		if index == -1 {
			assert.Nil(t, curr2.Random)
		} else {
			assert.Same(t, nodes[index], curr2.Random)
		}
		curr2 = curr2.Next
		i++
	}

}
