package leetcode

func rec(n int, openCount int, closeCount int, current string, result *[]string) {
	if openCount == n && closeCount == n {
		*result = append(*result, current)
		return
	}
	if openCount < n {
		rec(n, openCount+1, closeCount, current+"(", result)
	}
	if closeCount < openCount {
		rec(n, openCount, closeCount+1, current+")", result)
	}
}

func generateParenthesis(n int) []string {
	result := []string{}
	rec(n, 0, 0, "", &result)
	return result
}
