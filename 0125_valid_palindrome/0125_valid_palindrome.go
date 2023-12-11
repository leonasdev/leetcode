package leetcode

import (
	"strings"
	"unicode"
)

func isPalindrome(s string) bool {
	s = strings.ToLower(s)
	runes := []rune{}
	for _, ch := range s {
		if unicode.IsLetter(ch) || unicode.IsDigit(ch) {
			runes = append(runes, ch)
		}
	}

	lo, hi := 0, len(runes)-1
	for lo <= hi {
		if runes[lo] != runes[hi] {
			return false
		}
		lo++
		hi--
	}
	return true
}
