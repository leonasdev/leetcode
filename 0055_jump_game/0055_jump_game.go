package leetcode

import "fmt"

func canJump(nums []int) bool {
	maxStep := 0
	for index, val := range nums {
		if index > maxStep {
			return false
		}
		if index+val > maxStep {
			maxStep = index + val
		}
	}
	return true
}

func main() {
	nums := []int{2, 3, 1, 1, 4}
	fmt.Println(canJump(nums)) // false

	nums = []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums)) // false

	nums = []int{0}
	fmt.Println(canJump(nums)) // true

	nums = []int{1, 0}
	fmt.Println(canJump(nums)) // true

	nums = []int{2, 0, 0}
	fmt.Println(canJump(nums)) // true

	nums = []int{2, 2, 0, 1, 4} // true
	fmt.Println(canJump(nums))
}
