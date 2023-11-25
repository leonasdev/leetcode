package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntsToRandomList(t *testing.T) {
	nums := [][2]int{{1, 1}, {2, -1}}
	head := IntsToRandomList(nums)

	curr := head
	nodes := []*RandomListNode{}
	for i := 0; curr != nil; i++ {
		assert.Equal(t, nums[i][0], curr.Val)
		nodes = append(nodes, curr)
		curr = curr.Next
	}

	curr = head
	for i := 0; curr != nil; i++ {
		index := nums[i][1]
		if index == -1 {
			assert.Nil(t, curr.Random)
		} else {
			assert.Equal(t, nodes[index], curr.Random)
		}
		curr = curr.Next
	}
}
