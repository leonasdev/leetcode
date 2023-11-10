package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestJump(t *testing.T) {
	assert.Equal(t, 2, jump([]int{2, 3, 1, 1, 4}))
	assert.Equal(t, 2, jump([]int{1, 2, 3}))
	assert.Equal(t, 3, jump([]int{1, 1, 1, 1}))
	assert.Equal(t, 1, jump([]int{6, 1, 1, 1}))
	assert.Equal(t, 3, jump([]int{3, 2, 3, 1, 1, 1, 1}))
	assert.Equal(t, 1, jump([]int{7, 2, 3, 1, 1, 1, 1}))
	assert.Equal(t, 0, jump([]int{0}))
	assert.Equal(t, 1, jump([]int{8, 1}))
}
