package leetcode

import "github.com/leonasdev/leetcode/structure"

type Node = structure.RandomListNode

func copyRandomList(head *Node) *Node {
	dummy := &Node{}
	curr := head
	curr2 := dummy
	m := map[*Node]*Node{}

	for curr != nil {
		newNode := &Node{Val: curr.Val}
		m[curr] = newNode
		curr2.Next = newNode
		curr = curr.Next
		curr2 = curr2.Next
	}

	curr = head
	curr2 = dummy.Next
	for curr2 != nil {
		curr2.Random = m[curr.Random]
		curr = curr.Next
		curr2 = curr2.Next
	}

	return dummy.Next
}
