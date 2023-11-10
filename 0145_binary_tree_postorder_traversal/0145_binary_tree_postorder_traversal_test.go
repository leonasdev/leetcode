package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestPostorderTraversal(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	assert.Equal(t, []int{3, 2, 1}, postorderTraversal(root)) // [3,2,1]
	root = &TreeNode{
		Val: 1,
	}
	assert.Equal(t, []int{1}, postorderTraversal(root)) // [3,2,1]
	assert.Equal(t, []int{}, postorderTraversal(nil))   // [3,2,1]
}
