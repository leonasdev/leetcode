package main

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
	*res = append(*res, node.Val)
	traverse(node.Right, res)
}

func inorderTraversal(root *TreeNode) []int {
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
	fmt.Println(inorderTraversal(root)) // [1,3,2]
	root = &TreeNode{
		Val: 1,
	}
	fmt.Println(inorderTraversal(root)) // [1]
	fmt.Println(inorderTraversal(nil))  // []
}
