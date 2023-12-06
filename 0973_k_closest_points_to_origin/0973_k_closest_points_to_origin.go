package leetcode

import (
	"container/heap"
	"sort"
)

type Point struct {
	x        int
	y        int
	distance int
}

// implement heap.interface
type PriorityQueue []*Point

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Swap(i, j int)      { pq[j], pq[i] = pq[i], pq[j] }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].distance < pq[j].distance }
func (pq *PriorityQueue) Push(x any)        { *pq = append(*pq, x.(*Point)) }
func (pq *PriorityQueue) Pop() any {
	x := (*pq)[len(*pq)-1]
	(*pq) = (*pq)[:len(*pq)-1]
	return x
}

type MinHeap [][]int

func (h MinHeap) Len() int      { return len(h) }
func (h MinHeap) Swap(i, j int) { h[j], h[i] = h[i], h[j] }
func (h MinHeap) Less(i, j int) bool {
	return h[i][0]*h[i][0]+h[i][1]*h[i][1] < h[j][0]*h[j][0]+h[j][1]*h[j][1]
}
func (h *MinHeap) Push(x any) { *h = append(*h, x.([]int)) }
func (h *MinHeap) Pop() any {
	x := (*h)[len(*h)-1]
	(*h) = (*h)[:len(*h)-1]
	return x
}

func kClosestMinHeapUsingInit(points [][]int, k int) [][]int {
	pq := MinHeap(points)
	heap.Init(&pq)

	res := [][]int{}
	for i := 0; i < k; i++ {
		point := heap.Pop(&pq).([]int)
		res = append(res, point)
	}

	return res
}

func kClosestMinHeap(points [][]int, k int) [][]int {
	pq := &PriorityQueue{}
	for _, p := range points {
		point := &Point{
			x:        p[0],
			y:        p[1],
			distance: p[0]*p[0] + p[1]*p[1],
		}
		heap.Push(pq, point)
	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		point := heap.Pop(pq).(*Point)
		res = append(res, []int{point.x, point.y})
	}

	return res
}

func kClosestMaxHeap(points [][]int, K int) [][]int {
	pq := &PriorityQueue{}
	for _, p := range points {
		point := &Point{
			x:        p[0],
			y:        p[1],
			distance: -(p[0]*p[0] + p[1]*p[1]),
		}
		heap.Push(pq, point)
		if pq.Len() > K {
			heap.Pop(pq)
		}
	}

	res := [][]int{}
	for pq.Len() > 0 {
		point := heap.Pop(pq).(*Point)
		res = append(res, []int{point.x, point.y})
	}
	return res
}

func kClosestSort(points [][]int, k int) [][]int {
	sort.Slice(points, func(i, j int) bool {
		return points[i][0]*points[i][0]+points[i][1]*points[i][1] <
			points[j][0]*points[j][0]+points[j][1]*points[j][1]
	})
	return points[:k]
}

func kClosestMyHeap(points [][]int, k int) [][]int {
	h := &MyHeap{}
	for _, p := range points {
		h.Push(Point{
			x:        p[0],
			y:        p[1],
			distance: p[0]*p[0] + p[1]*p[1],
		})
	}

	res := [][]int{}
	for i := 0; i < k; i++ {
		p := h.Pop()
		res = append(res, []int{p.x, p.y})
	}

	return res
}
