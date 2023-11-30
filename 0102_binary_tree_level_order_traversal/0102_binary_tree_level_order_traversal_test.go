package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestLevelOrder(t *testing.T) {
	t.Run("1", func(t *testing.T) {
		nums := []int{3, 9, 20, structure.NULL, structure.NULL, 15, 7}
		want := [][]int{{3}, {9, 20}, {15, 7}}
		assert.Equal(t, want, levelOrder(structure.IntsToTreeNode(nums)))
	})

	t.Run("2", func(t *testing.T) {
		nums := []int{1}
		want := [][]int{{1}}
		assert.Equal(t, want, levelOrder(structure.IntsToTreeNode(nums)))
	})

	t.Run("1", func(t *testing.T) {
		want := [][]int{}
		assert.Equal(t, want, levelOrder(structure.IntsToTreeNode(nil)))
	})
}
