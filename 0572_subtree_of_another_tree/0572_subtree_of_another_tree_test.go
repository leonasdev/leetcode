package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestIsSubtree(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		root := structure.IntsToTreeNode([]int{3, 4, 5, 1, 2})
		subRoot := structure.IntsToTreeNode([]int{4, 1, 2})
		assert.True(t, isSubtree(root, subRoot))
	})

	t.Run("2", func(t *testing.T) {
		root := structure.IntsToTreeNode([]int{3, 4, 5, 1, 2, structure.NULL, structure.NULL, structure.NULL, structure.NULL, 0})
		subRoot := structure.IntsToTreeNode([]int{4, 1, 2})
		assert.False(t, isSubtree(root, subRoot))
	})
}
