package structure

type RandomListNode struct {
	Val    int
	Next   *RandomListNode
	Random *RandomListNode
}

// ListToInts convert Node to []int
func RandomListToInts(head *RandomListNode) []int {
	res := []int{}

	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}

	return res
}

// IntsToRandomList convert [][2]int to Node
// Each node is represented as a pair of [val, random_index] where:
//
// val: an integer representing Node.val
//
// random_index: the index of the node (range from 0 to n-1) that the random pointer points to,
// or -1 if it does not point to any node.
//
// For example, if the given linked list is [[1,1],[2,-1]], it means that there are two nodes
// in the linked list, where the first node (val = 1) points to the second node and the second
// node (val = 2) point to nil.
func IntsToRandomList(nums [][2]int) *RandomListNode {
	if len(nums) == 0 {
		return nil
	}

	m := map[int]*RandomListNode{}

	dummyHead := &RandomListNode{}
	curr := dummyHead
	i := 0
	for _, num := range nums {
		curr.Next = &RandomListNode{Val: num[0]}
		curr = curr.Next
		m[i] = curr
		i++
	}

	curr = dummyHead.Next
	for _, num := range nums {
		curr.Random = m[num[1]]
		curr = curr.Next
	}

	return dummyHead.Next
}
