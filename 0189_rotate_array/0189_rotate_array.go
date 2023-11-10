package leetcode

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
