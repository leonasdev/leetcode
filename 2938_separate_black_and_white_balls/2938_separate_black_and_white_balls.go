package leetcode

func minimumSteps(s string) int64 {
	count, res := 0, 0
	for _, c := range s {
		if c == '1' {
			count++
		} else {
			res += count
		}
	}
	return int64(res)
}
