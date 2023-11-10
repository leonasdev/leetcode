package leetcode

func maxInt(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func jump(nums []int) int {
	farest := 0
	end := 0
	jumps := 0
	for i := 0; i < len(nums)-1; i++ {
		farest = maxInt(farest, i+nums[i])

		if i == end {
			jumps++
			end = farest
		}
	}
	return jumps
}
