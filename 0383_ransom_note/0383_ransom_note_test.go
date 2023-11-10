package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCanConstruct(t *testing.T) {
	assert.Equal(t, false, canConstruct("a", "b"))   // false
	assert.Equal(t, false, canConstruct("aa", "ab")) // false
	assert.Equal(t, true, canConstruct("aa", "aab")) // true
}
