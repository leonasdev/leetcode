package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPreorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	assert.Equal(t, []int{1, 2, 3}, preorderTraversal(root))
	root = &TreeNode{
		Val: 1,
	}
	assert.Equal(t, []int{1}, preorderTraversal(root))
	assert.Equal(t, []int{}, preorderTraversal(nil))
}
