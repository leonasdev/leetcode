package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindContentChildren(t *testing.T) {
	assert.Equal(t, 1, findContentChildren([]int{1, 2, 3}, []int{1, 1}))
	assert.Equal(t, 2, findContentChildren([]int{1, 2}, []int{1, 2, 3}))
	assert.Equal(t, 2, findContentChildren([]int{10, 9, 8, 7}, []int{5, 6, 7, 8}))
}
