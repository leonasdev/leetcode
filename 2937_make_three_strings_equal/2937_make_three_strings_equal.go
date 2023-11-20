package leetcode

func findMinimumOperations(s1 string, s2 string, s3 string) int {
	minLen := min(len(s1), len(s2), len(s3))

	i := 0
	for ; i < minLen; i++ {
		if !(s1[i] == s2[i] && s2[i] == s3[i]) {
			break
		}
	}
	if i == 0 {
		return -1
	}
	return len(s1) + len(s2) + len(s3) - i*3
}
