package leetcode

import "github.com/leonasdev/leetcode/structure"

type TreeNode = structure.TreeNode

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	res := [][]int{}
	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelVals := []int{}
		qLen := len(queue)
		for i := 0; i < qLen; i++ {
			// dequeue
			node := queue[0]
			queue = queue[1:]

			levelVals = append(levelVals, node.Val)

			if node.Left != nil {
				queue = append(queue, node.Left)
			}

			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		res = append(res, levelVals)
	}

	return res
}
