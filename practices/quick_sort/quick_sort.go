package leetcode

func partition(arr []int, lo int, hi int) int {
	pivot := arr[hi]
	index := lo
	for i := lo; i < hi; i++ {
		if arr[i] <= pivot {
			arr[i], arr[index] = arr[index], arr[i]
			index++
		}
	}
	arr[index], arr[hi] = arr[hi], arr[index]
	return index
}

func qs(arr []int, lo int, hi int) {
	if lo >= hi {
		return
	}
	index := partition(arr, lo, hi)
	qs(arr, lo, index-1)
	qs(arr, index+1, hi)
}

func quickSort(arr []int) {
	qs(arr, 0, len(arr)-1)
}
