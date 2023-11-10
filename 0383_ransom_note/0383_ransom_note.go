package leetcode

func canConstruct(ransomNote string, magazine string) bool {
	letterToCount := map[rune]int{}
	for _, c := range magazine {
		letterToCount[c]++
	}
	for _, c := range ransomNote {
		letterToCount[c]--
		if letterToCount[c] < 0 {
			return false
		}
	}
	return true
}
