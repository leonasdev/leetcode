package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestDiameterOfBinaryTree(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		root := structure.IntsToTreeNode([]int{1, 2, 3, 4, 5})
		assert.Equal(t, 3, diameterOfBinaryTree(root))
	})

	t.Run("2", func(t *testing.T) {
		root := structure.IntsToTreeNode([]int{1, 2})
		assert.Equal(t, 1, diameterOfBinaryTree(root))
	})
}
