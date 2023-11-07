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
	traverse(node.Left, res)
	traverse(node.Right, res)
	*res = append(*res, node.Val)
}

func postorderTraversal(root *TreeNode) []int {
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
	fmt.Println(postorderTraversal(root)) // [3,2,1]
	root = &TreeNode{
		Val: 1,
	}
	fmt.Println(postorderTraversal(root)) // [1]
	fmt.Println(postorderTraversal(nil))  // []
}
