package leetcode

import "slices"

func characterReplacement(s string, k int) int {
	var lo, hi, longest int
	freq := make([]int, 26)

	for hi < len(s) {
		freq[s[hi]-'A']++
		if (hi-lo+1)-slices.Max(freq) > k {
			freq[s[lo]-'A']--
			lo++
		}
		longest = max(longest, hi-lo+1)
		hi++
	}

	return longest
}
