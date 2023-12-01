package leetcode

import "github.com/leonasdev/leetcode/structure"

type TreeNode = structure.TreeNode

func abs(a int) int {
	if a < 0 {
		return a * -1
	}
	return a
}

func balancedAndHeight(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	}

	lBalanced, lHeight := balancedAndHeight(root.Left)
	rBalanced, rHeight := balancedAndHeight(root.Right)

	if !lBalanced || !rBalanced {
		return false, 0
	}

	if abs(lHeight-rHeight) > 1 {
		return false, 0
	}

	return true, max(lHeight, rHeight) + 1
}

func isBalanced(root *TreeNode) bool {
	res, _ := balancedAndHeight(root)
	return res
}
