package leetcode

func lengthOfLongestSubstring(s string) int {
	lo, hi, maxLen := 0, 0, 0
	m := map[byte]int{}
	for hi < len(s) {
		m[s[hi]]++
		for m[s[hi]] > 1 {
			m[s[lo]]--
			lo++
		}
		maxLen = max(hi-lo+1, maxLen)
		hi++
	}
	return maxLen
}
