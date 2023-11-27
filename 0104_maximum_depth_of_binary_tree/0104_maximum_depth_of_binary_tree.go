package leetcode

import "github.com/leonasdev/leetcode/structure"

type TreeNode = structure.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(maxDepth(root.Left), maxDepth(root.Right)) + 1
}
