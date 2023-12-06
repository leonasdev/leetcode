package leetcode

import (
	"math/rand"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKClosestMinHeap(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMinHeap([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMinHeap([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMinHeap([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMinHeap([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMinHeap([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMinHeap([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMinHeap([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMinHeap([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
}

func TestKClosestMinHeapUsingInit(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMinHeapUsingInit([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMinHeapUsingInit([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMinHeapUsingInit([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMinHeapUsingInit([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMinHeapUsingInit([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMinHeapUsingInit([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMinHeapUsingInit([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMinHeapUsingInit([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
}

func TestKClosestMaxHeap(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMaxHeap([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMaxHeap([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMaxHeap([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMaxHeap([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMaxHeap([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMaxHeap([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMaxHeap([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMaxHeap([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
}

func TestKClosestSort(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestSort([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestSort([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestSort([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestSort([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestSort([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestSort([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestSort([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestSort([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
}

func TestMyHeap(t *testing.T) {
	h := &MyHeap{}
	h.Push(Point{1, 2, 3})
	h.Push(Point{2, 3, 4})
	h.Push(Point{3, 4, 5})
	assert.Equal(t, 3, h.Len())
	assert.Equal(t, Point{1, 2, 3}, h.Pop())
	assert.Equal(t, Point{2, 3, 4}, h.Pop())
	assert.Equal(t, Point{3, 4, 5}, h.Pop())
	assert.Equal(t, 0, h.Len())
}

func TestKClosestMyHeap(t *testing.T) {
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMyHeap([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMyHeap([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMyHeap([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMyHeap([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
	assert.ElementsMatch(t, [][]int{{-2, 2}}, kClosestMyHeap([][]int{{1, 3}, {-2, 2}}, 1))
	assert.ElementsMatch(t, [][]int{{3, 3}, {-2, 4}}, kClosestMyHeap([][]int{{3, 3}, {5, -1}, {-2, 4}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}}, kClosestMyHeap([][]int{{1, 3}, {-2, 2}}, 2))
	assert.ElementsMatch(t, [][]int{{-2, 2}, {1, 3}, {3, 3}}, kClosestMyHeap([][]int{{1, 3}, {-2, 2}, {3, 3}}, 3))
}

var nums [][]int

func setup() {
	rand.New(rand.NewSource(99))
	nums = [][]int{}
	for i := 0; i < 100000; i++ {
		nums = append(nums, []int{rand.Intn(1000), rand.Intn(1000)})
	}
}

func BenchmarkKClosestMinHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosestMinHeap(nums, len(nums)/2)
	}
}

func BenchmarkKClosestMaxHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosestMaxHeap(nums, len(nums)/2)
	}
}

func BenchmarkKClosestSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosestSort(nums, len(nums)/2)
	}
}

func BenchmarkKClosestMinHeapUsingInit(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosestMinHeapUsingInit(nums, len(nums)/2)
	}
}

func BenchmarkKClosestMyHeap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		kClosestMyHeap(nums, len(nums)/2)
	}
}

func TestMain(m *testing.M) {
	setup()
	os.Exit(m.Run())
}
