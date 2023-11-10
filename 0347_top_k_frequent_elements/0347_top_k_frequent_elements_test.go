package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTopKFrequent(t *testing.T) {
	nums := []int{1, 1, 2, 2, 1, 3}
	assert.Equal(t, []int{1, 2}, topKFrequent(nums, 2)) // [1, 2]

	nums = []int{1}
	assert.Equal(t, []int{1}, topKFrequent(nums, 1)) // [1]
}
