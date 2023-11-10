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
	traverse(node.Left, res)
	traverse(node.Right, res)
	*res = append(*res, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
	res := []int{}
	traverse(root, &res)
	return res
}
