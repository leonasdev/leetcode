package leetcode

import "github.com/leonasdev/leetcode/structure"

type TreeNode = structure.TreeNode

func invert(root *TreeNode) {
	if root == nil {
		return
	}
	invert(root.Left)
	invert(root.Right)
	root.Left, root.Right = root.Right, root.Left
}

func invertTree(root *TreeNode) *TreeNode {
	invert(root)
	return root
}
