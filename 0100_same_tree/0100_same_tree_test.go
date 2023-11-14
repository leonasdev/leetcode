package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsSameTree1(t *testing.T) {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 3,
		},
	}
	assert.Equal(t, true, isSameTree(p, q))
}

func TestIsSameTree2(t *testing.T) {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
	}
	q := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 3,
		},
	}
	assert.Equal(t, false, isSameTree(p, q))
}

func TestIsSameTree3(t *testing.T) {
	p := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
		},
		Right: &TreeNode{
			Val: 1,
		},
	}
	q := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 1,
		},
		Right: &TreeNode{
			Val: 2,
		},
	}
	assert.Equal(t, false, isSameTree(p, q))
}
