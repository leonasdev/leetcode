package leetcode

import "container/heap"

// Implement heap.interface
type MaxHeap []int

func (h MaxHeap) Len() int           { return len(h) }
func (h MaxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h MaxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *MaxHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *MaxHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

func lastStoneWeight(stones []int) int {
	h := &MaxHeap{}

	for _, s := range stones {
		heap.Push(h, s)
	}

	for h.Len() > 1 {
		y := heap.Pop(h)
		x := heap.Pop(h)
		newWeight := y.(int) - x.(int)
		if newWeight != 0 {
			heap.Push(h, newWeight)
		}
	}

	if h.Len() == 0 {
		return 0
	}

	return (*h)[0]
}
