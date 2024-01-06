package leetcode

func search(nums []int, target int) int {
	lo, hi := 0, len(nums)-1

	for lo <= hi {
		mid := lo + (hi-lo)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] >= nums[lo] { // left portion
			if target > nums[mid] || target < nums[lo] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		} else { // right portion
			if target < nums[mid] || target > nums[hi] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		}
	}

	return -1
}
