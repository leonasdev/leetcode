package leetcode

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func traverse(node *TreeNode, res *[]int) {
	if node == nil {
		return
	}
	*res = append(*res, node.Val)
	traverse(node.Left, res)
	traverse(node.Right, res)
}

func preorderTraversal(root *TreeNode) []int {
	res := []int{}
	traverse(root, &res)
	return res
}
