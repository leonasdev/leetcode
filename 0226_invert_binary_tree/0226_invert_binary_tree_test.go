package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestInvertTree(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		expect := structure.IntsToTreeNode([]int{4, 7, 2, 9, 6, 3, 1})
		actual := invertTree(structure.IntsToTreeNode([]int{4, 2, 7, 1, 3, 6, 9}))
		assert.True(t, structure.SameTree(expect, actual))
	})

	t.Run("2", func(t *testing.T) {
		var expect *TreeNode
		actual := invertTree(nil)
		assert.True(t, structure.SameTree(expect, actual))
	})
}
