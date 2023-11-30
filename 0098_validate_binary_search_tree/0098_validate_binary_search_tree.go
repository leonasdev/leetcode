package leetcode

import (
	"math"

	"github.com/leonasdev/leetcode/structure"
)

type TreeNode = structure.TreeNode

func recHelper(root *TreeNode, minVal int, maxVal int) bool {
	if root == nil {
		return true
	}

	if root.Val <= minVal || root.Val >= maxVal {
		return false
	}

	return recHelper(root.Left, minVal, root.Val) && recHelper(root.Right, root.Val, maxVal)
}

func isValidBST(root *TreeNode) bool {
	return recHelper(root, math.MinInt, math.MaxInt)
}
