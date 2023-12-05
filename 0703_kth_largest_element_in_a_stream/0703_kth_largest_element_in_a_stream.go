package leetcode

import "container/heap"

// Implement heap.interface
type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *IntHeap) Push(x any)        { *h = append(*h, x.(int)) }
func (h *IntHeap) Pop() any {
	x := (*h)[len(*h)-1]
	*h = (*h)[:len(*h)-1]
	return x
}

type KthLargest struct {
	heap *IntHeap
	k    int
}

func Constructor(k int, nums []int) KthLargest {
	obj := KthLargest{
		heap: &IntHeap{},
		k:    k,
	}

	for _, n := range nums {
		obj.Add(n)
	}

	return obj
}

func (kth *KthLargest) Add(val int) int {
	heap.Push(kth.heap, val)
	for kth.heap.Len() > kth.k {
		heap.Pop(kth.heap)
	}
	return (*kth.heap)[0]
}
