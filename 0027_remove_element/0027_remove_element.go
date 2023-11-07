package leetcode

func removeElement(nums []int, val int) int {
	result := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != val {
			nums[result] = nums[i]
			result++
		}
	}
	return result
}

func main() {
	nums1 := []int{3, 2, 2, 3}
	val1 := 3
	println(removeElement(nums1, val1))

	nums2 := []int{0, 1, 2, 2, 3, 0, 4, 2}
	val2 := 2
	println(removeElement(nums2, val2))
}
