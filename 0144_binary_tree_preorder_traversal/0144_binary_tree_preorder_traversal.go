package leetcode

import "fmt"

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

func main() {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}
	fmt.Println(preorderTraversal(root)) // [1,2,3]
	root = &TreeNode{
		Val: 1,
	}
	fmt.Println(preorderTraversal(root)) // [1]
	fmt.Println(preorderTraversal(nil))  // []
}
