package leetcode

func checkInclusion(s1 string, s2 string) bool {
	if len(s1) > len(s2) {
		return false
	}

	freq := [26]int{}
	match := [26]int{}

	for i := range s1 {
		match[s1[i]-'a']++
		freq[s2[i]-'a']++
	}

	lo, hi := 0, len(s1)-1
	for {
		if match == freq {
			return true
		}
		hi++
		if hi >= len(s2) {
			break
		}
		freq[s2[hi]-'a']++
		freq[s2[lo]-'a']--
		lo++
	}

	return false
}
