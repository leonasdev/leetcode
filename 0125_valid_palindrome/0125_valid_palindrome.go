package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	alphaSlice := []rune{}
	for _, ch := range s {
		if unicode.IsDigit(ch) || (ch >= 'A' && ch <= 'Z') || (ch >= 'a' && ch <= 'z') {
			alphaSlice = append(alphaSlice, ch)
		}
	}

	for i, j := 0, len(alphaSlice)-1; i < len(alphaSlice); i, j = i+1, j-1 {
		if alphaSlice[i] != alphaSlice[j] {
			return false
		}
	}
	return true
}
