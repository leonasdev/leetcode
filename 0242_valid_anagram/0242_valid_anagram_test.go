package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsAnagram(t *testing.T) {
	assert.Equal(t, true, isAnagram("anagram", "nagaram"))
	assert.Equal(t, false, isAnagram("rat", "car"))
	assert.Equal(t, true, isAnagram("你好不好", "好不好你"))
	assert.Equal(t, false, isAnagram("你好不", "好不好你"))
	assert.Equal(t, true, isAnagram("你好不", "不好你"))
	assert.Equal(t, false, isAnagram("你好不", "不你"))
}
