package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestIsBalanced(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		root := structure.IntsToTreeNode([]int{3, 9, 20, structure.NULL, structure.NULL, 15, 7})
		assert.True(t, isBalanced(root))
	})

	t.Run("2", func(t *testing.T) {
		root := structure.IntsToTreeNode([]int{1, 2, 2, 3, 3, structure.NULL, structure.NULL, 4, 4})
		assert.False(t, isBalanced(root))
	})
}
