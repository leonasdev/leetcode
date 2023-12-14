package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxArea(t *testing.T) {
	assert.Equal(t, 49, maxArea([]int{1, 8, 6, 2, 5, 4, 8, 3, 7}))
	assert.Equal(t, 1, maxArea([]int{1, 1}))
	assert.Equal(t, 0, maxArea([]int{1, 0}))
	assert.Equal(t, 0, maxArea([]int{}))
	assert.Equal(t, 1, maxArea([]int{1, 999}))
}
