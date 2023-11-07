package leetcode

import (
	"fmt"
)

func search(nums []int, target int) int {
	for l, r := 0, len(nums); l < r; {
		mid := l + (r-l)/2 // prevent overflow when the sum of r and l is greater than 2^32
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			r = mid
		} else {
			l = mid + 1
		}
	}
	return -1
}

func main() {
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 9)) // 4
	fmt.Println(search([]int{-1, 0, 3, 5, 9, 12}, 2)) // -1
}
