package leetcode

import (
	"fmt"
)

func containsNearbyDuplicate(nums []int, k int) bool {
	record := map[int]int{}
	for i, v := range nums {
		if j, ok := record[v]; ok && i-j <= k {
			return true
		}
		record[v] = i
	}
	return false
}

func main() {
	nums := []int{1, 2, 3, 1}
	k := 3
	fmt.Println(containsNearbyDuplicate(nums, k)) // true
	nums = []int{1, 0, 1, 1}
	k = 1
	fmt.Println(containsNearbyDuplicate(nums, k)) // true
	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	fmt.Println(containsNearbyDuplicate(nums, k)) // false
}
