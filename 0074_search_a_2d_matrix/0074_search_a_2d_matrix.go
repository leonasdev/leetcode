package leetcode

func search(nums []int, target int) bool {
	lo, hi := 0, len(nums)
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if target < nums[mid] {
			hi = mid - 1
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			return true
		}
	}
	return false
}

func searchMatrix(matrix [][]int, target int) bool {
	lo, hi := 0, len(matrix)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		nums := matrix[mid]
		if target == nums[0] || target == nums[len(nums)-1] {
			return true
		} else if target > nums[0] && target < nums[len(nums)-1] {
			return search(nums, target)
		} else if target > nums[len(nums)-1] {
			lo = mid + 1
		} else if target < nums[0] {
			hi = mid - 1
		}
	}
	return false
}
