package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxDepth(t *testing.T) {
	root := &TreeNode{
		Val: 3,
		Left: &TreeNode{
			Val: 9,
		},
		Right: &TreeNode{
			Val: 20,
			Left: &TreeNode{
				Val: 15,
			},
			Right: &TreeNode{
				Val: 7,
			},
		},
	}
	assert.Equal(t, 3, maxDepth(root))
	root = &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
		},
	}
	assert.Equal(t, 2, maxDepth(root))
	assert.Equal(t, 0, maxDepth(nil))
}
