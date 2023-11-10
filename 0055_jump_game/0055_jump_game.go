package leetcode

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
