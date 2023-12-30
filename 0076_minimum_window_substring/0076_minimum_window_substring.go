package leetcode

func minWindow(s string, t string) string {
	if len(t) > len(s) {
		return ""
	}

	match := [128]int{}
	set := map[byte]int{}
	for i := 0; i < len(t); i++ {
		match[t[i]]++
		set[t[i]]++
	}

	curr := [128]int{}
	lo, minStr, needSatisfies := 0, "", len(set)
	for hi := 0; hi < len(s); hi++ {
		curr[s[hi]]++
		if match[s[hi]] == curr[s[hi]] {
			needSatisfies--
		}
		for needSatisfies == 0 {
			if minStr == "" || (hi-lo+1) < len(minStr) {
				minStr = s[lo : hi+1]
			}
			curr[s[lo]]--
			if match[s[lo]] > curr[s[lo]] {
				needSatisfies++
			}
			lo++
		}
	}

	return minStr
}
