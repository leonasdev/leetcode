package leetcode

func trap(height []int) int {
	res := 0
	lo, hi := 1, len(height)-2
	maxLo, maxHi := height[0], height[len(height)-1]

	for lo <= hi {
		maxLo = max(maxLo, height[lo-1])
		maxHi = max(maxHi, height[hi+1])
		trap := 0

		if maxLo <= maxHi {
			trap = maxLo - height[lo]
			lo++
		} else {
			trap = maxHi - height[hi]
			hi--
		}

		if trap > 0 {
			res += trap
		}
	}

	return res
}
