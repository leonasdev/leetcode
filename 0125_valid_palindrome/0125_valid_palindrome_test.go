package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsPalindrome(t *testing.T) {
	assert.Equal(t, true, isPalindrome("A man, a plan, a canal: Panama"))
	assert.Equal(t, false, isPalindrome("flllfds"))
	assert.Equal(t, true, isPalindrome("ab_a"))
}
