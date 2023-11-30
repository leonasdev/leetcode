package leetcode

import (
	"testing"

	"github.com/leonasdev/leetcode/structure"
	"github.com/stretchr/testify/assert"
)

func TestIsValidBST(t *testing.T) {
	assert.True(t, isValidBST(structure.IntsToTreeNode([]int{2, 1, 3})))
	assert.False(t, isValidBST(structure.IntsToTreeNode([]int{5, 1, 4, structure.NULL, structure.NULL, 3, 6})))
	assert.False(t, isValidBST(structure.IntsToTreeNode([]int{2, 2, 2})))
	assert.False(t, isValidBST(structure.IntsToTreeNode([]int{5, 1, 8, 7, 9})))
}
