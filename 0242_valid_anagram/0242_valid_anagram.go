package leetcode

func isAnagram(s string, t string) bool {
	letterToCount := map[rune]int{}
	for _, v := range s {
		letterToCount[v]++
	}
	for _, v := range t {
		letterToCount[v]--
	}
	for _, v := range letterToCount {
		if v != 0 {
			return false
		}
	}
	return true
}
