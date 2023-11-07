package leetcode

import "fmt"

func rotate(nums []int, k int) {
	k = k % len(nums)
	k = len(nums) - k
	var tmp []int
	tmp = append(tmp, nums[k:]...)
	tmp = append(tmp, nums[:k]...)
	for i := 0; i < len(nums); i++ {
		nums[i] = tmp[i]
	}
}

func main() {
	nums := []int{1, 2, 3, 4, 5, 6, 7}
	k := 3
	rotate(nums, k)
	fmt.Println(nums)

	nums = []int{-1, -100, 3, 99}
	k = 2
	rotate(nums, k)
	fmt.Println(nums)
}
