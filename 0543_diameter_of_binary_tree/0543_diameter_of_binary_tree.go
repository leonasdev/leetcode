package leetcode

import "github.com/leonasdev/leetcode/structure"

type TreeNode = structure.TreeNode

func diameterOfBinaryTree(root *TreeNode) int {
	maxDiameter := 0
	var heightOfTree func(*TreeNode) int

	heightOfTree = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		lHeight := heightOfTree(root.Left)
		rHeight := heightOfTree(root.Right)

		maxDiameter = max(lHeight+rHeight, maxDiameter)

		return max(lHeight, rHeight) + 1
	}

	heightOfTree(root)

	return maxDiameter
}
