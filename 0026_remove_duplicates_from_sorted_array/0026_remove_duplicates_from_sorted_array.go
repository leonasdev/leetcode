package leetcode

import "fmt"

func removeDuplicates(nums []int) int {
	i := 0
	for j := range nums {
		if nums[i] != nums[j] {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}

func main() {
	nums := []int{1, 1, 2}
	result := removeDuplicates(nums)
	fmt.Println(nums, result)

	nums = []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	result = removeDuplicates(nums)
	fmt.Println(nums, result)
}
