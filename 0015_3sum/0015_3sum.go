package leetcode

import "sort"

func threeSum(nums []int) [][]int {
	res := [][]int{}
	sort.Ints(nums)
	for i, n := range nums {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		lo, hi := i+1, len(nums)-1
		for lo < hi {
			sum := nums[lo] + nums[hi] + n
			if sum == 0 {
				res = append(res, []int{nums[lo], nums[hi], n})
				lo++
				for lo < hi && nums[lo] == nums[lo-1] {
					lo++
				}
			} else if sum < 0 {
				lo++
			} else if sum > 0 {
				hi--
			}
		}
	}

	return res
}
