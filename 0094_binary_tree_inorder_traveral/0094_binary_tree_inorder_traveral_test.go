package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestInorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	assert.Equal(t, []int{1, 3, 2}, inorderTraversal(root))

	root = &TreeNode{
		Val: 1,
	}
	assert.Equal(t, []int{1}, inorderTraversal(root))
}
