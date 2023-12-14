package leetcode

func maxArea(height []int) int {
	lo, hi := 0, len(height)-1
	maxAmount := 0

	for lo < hi {
		minHeight := min(height[lo], height[hi])
		maxAmount = max(maxAmount, minHeight*(hi-lo))
		if height[lo] < height[hi] {
			lo++
		} else {
			hi--
		}
	}

	return maxAmount
}
