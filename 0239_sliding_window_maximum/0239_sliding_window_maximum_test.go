package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxSlidingWindow(t *testing.T) {
	assert.Equal(t, []int{3, 3, 5, 5, 6, 7}, maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3))
	assert.Equal(t, []int{1}, maxSlidingWindow([]int{1}, 1))
	assert.Equal(t, []int{1, -1}, maxSlidingWindow([]int{1, -1}, 1))
	assert.Equal(t, []int{3, 3, 2, 5}, maxSlidingWindow([]int{1, 3, 1, 2, 0, 5}, 3))
}
