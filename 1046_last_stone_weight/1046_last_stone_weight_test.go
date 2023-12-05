package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLastStoneWeight(t *testing.T) {
	assert.Equal(t, 1, lastStoneWeight([]int{2, 7, 4, 1, 8, 1}))
	assert.Equal(t, 2, lastStoneWeight([]int{1, 3}))
	assert.Equal(t, 2, lastStoneWeight([]int{3, 7, 2}))
	assert.Equal(t, 1, lastStoneWeight([]int{3, 7, 2, 1}))
	assert.Equal(t, 0, lastStoneWeight([]int{3, 7, 2, 1, 1}))
}
