package structure

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIntsToTreeNode(t *testing.T) {
	expectedTree := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 4,
			},
			Right: &TreeNode{
				Val: 5,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val: 6,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}

	nums := []int{1, 2, 3, 4, 5, 6, 7}
	assert.True(t, SameTree(IntsToTreeNode(nums), expectedTree))
}
