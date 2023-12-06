package leetcode

type MyHeap struct {
	elements []Point
	size     int
}

func (h *MyHeap) heapifyUp(index int) {
	for index > 0 && h.elements[(index-1)/2].distance > h.elements[index].distance {
		h.elements[(index-1)/2], h.elements[index] = h.elements[index], h.elements[(index-1)/2]
		index = (index - 1) / 2
	}
}

func (h *MyHeap) heapifyDown(index int) {
	for index < h.size {
		left, right := index*2+1, index*2+2
		smallest := index
		if left < h.size && h.elements[left].distance < h.elements[smallest].distance {
			smallest = left
		}
		if right < h.size && h.elements[right].distance < h.elements[smallest].distance {
			smallest = right
		}
		if smallest == index {
			break
		}
		h.elements[index], h.elements[smallest] = h.elements[smallest], h.elements[index]
		index = smallest
	}
}

func (h *MyHeap) Push(p Point) {
	h.elements = append(h.elements, p)
	h.size++
	h.heapifyUp(h.size - 1)
}

func (h *MyHeap) Pop() Point {
	p := h.elements[0]
	h.elements[0] = h.elements[h.size-1]
	h.elements = h.elements[:h.size-1]
	h.size--
	h.heapifyDown(0)
	return p
}

func (h *MyHeap) Len() int {
	return h.size
}
