package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGenerateParentesis(t *testing.T) {
	assert.Equal(t, generateParenthesis(3), []string{"((()))", "(()())", "(())()", "()(())", "()()()"})
}
