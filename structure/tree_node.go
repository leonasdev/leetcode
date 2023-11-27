package structure

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// TreeToInts convert TreeNode to []int
func TreeToInts(root *TreeNode) []int {
	res := []int{}

	var preorder func(root *TreeNode)
	preorder = func(root *TreeNode) {
		if root == nil {
			return
		}

		res = append(res, root.Val)
		preorder(root.Left)
		preorder(root.Right)
	}

	preorder(root)

	return res
}

// IntsToTree convert []int to TreeNode
var NULL = math.MinInt

// IntsToTreeNode convert []int to TreeNode
func IntsToTreeNode(nums []int) *TreeNode {
	n := len(nums)

	if n == 0 {
		return nil
	}

	root := &TreeNode{Val: nums[0]}
	queue := []*TreeNode{}
	queue = append(queue, root)

	i := 1
	for i < n {
		node := queue[0]
		queue = queue[1:]

		if i < n && nums[i] != NULL {
			node.Left = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Left)
		}
		i++

		if i < n && nums[i] != NULL {
			node.Right = &TreeNode{Val: nums[i]}
			queue = append(queue, node.Right)
		}
		i++
	}

	return root
}

func SameTree(t1, t2 *TreeNode) bool {
	if t1 == nil && t2 == nil {
		return true
	}

	if t1 == nil || t2 == nil {
		return false
	}

	if t1.Val != t2.Val {
		return false
	}

	return SameTree(t1.Left, t2.Left) && SameTree(t1.Right, t2.Right)
}
