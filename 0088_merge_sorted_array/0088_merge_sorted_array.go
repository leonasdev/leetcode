package leetcode

import (
	"fmt"
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	for i := 0; i < n; i++ {
		nums1[m+i] = nums2[i]
	}
	sort.Ints(nums1)
}

func main() {
	// merge the array
	// nums1 = [1,2,3,0,0,0], m = 3
	// nums2 = [2,5,6],       n = 3
	// Output: [1,2,2,3,5,6]
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	// merge another array
	// nums1 = [1], m = 1
	// nums2 = [],  n = 0
	// Output: [1]
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	// merge another array
	// nums1 = [0], m = 0
	// nums2 = [1],  n = 1
	// Output: [1]
	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

	// merge another array
	// nums1 = [4,0,0,0,0,0], m = 1
	// nums2 = [1,2,3,5,6],  n = 5
	// Output: [1,2,3,4,5,6]
	nums1 = []int{4, 0, 0, 0, 0, 0}
	m = 1
	nums2 = []int{1, 2, 3, 5, 6}
	n = 5
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)
}
