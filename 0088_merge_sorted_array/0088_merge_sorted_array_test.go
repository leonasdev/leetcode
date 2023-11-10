package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMerge(t *testing.T) {
	// merge the array
	// nums1 = [1,2,3,0,0,0], m = 3
	// nums2 = [2,5,6],       n = 3
	// Output: [1,2,2,3,5,6]
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)

	assert.Equal(t, []int{1, 2, 2, 3, 5, 6}, nums1)

	// merge another array
	// nums1 = [1], m = 1
	// nums2 = [],  n = 0
	// Output: [1]
	nums1 = []int{1}
	m = 1
	nums2 = []int{}
	n = 0
	merge(nums1, m, nums2, n)

	assert.Equal(t, []int{1}, nums1)

	// merge another array
	// nums1 = [0], m = 0
	// nums2 = [1],  n = 1
	// Output: [1]
	nums1 = []int{0}
	m = 0
	nums2 = []int{1}
	n = 1
	merge(nums1, m, nums2, n)

	assert.Equal(t, []int{1}, nums1)

	// merge another array
	// nums1 = [4,0,0,0,0,0], m = 1
	// nums2 = [1,2,3,5,6],  n = 5
	// Output: [1,2,3,4,5,6]
	nums1 = []int{4, 0, 0, 0, 0, 0}
	m = 1
	nums2 = []int{1, 2, 3, 5, 6}
	n = 5
	merge(nums1, m, nums2, n)

	assert.Equal(t, []int{1, 2, 3, 4, 5, 6}, nums1)
}
