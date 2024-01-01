package leetcode

func maxSlidingWindow(nums []int, k int) []int {
	deque := []int{} // store indices
	res := []int{}
	lo, hi := 0, 0

	for hi < len(nums) {
		for len(deque) > 0 && nums[deque[len(deque)-1]] < nums[hi] {
			deque = deque[:len(deque)-1]
		}

		deque = append(deque, hi)

		if hi >= k-1 {
			res = append(res, nums[deque[0]])
			lo++
			if lo > deque[0] {
				deque = deque[1:]
			}
		}

		hi++
	}

	return res
}
