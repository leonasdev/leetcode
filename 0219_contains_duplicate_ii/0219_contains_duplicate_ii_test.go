package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContainsNearbyDuplicate(t *testing.T) {
	nums := []int{1, 2, 3, 1}
	k := 3
	assert.Equal(t, true, containsNearbyDuplicate(nums, k))

	nums = []int{1, 0, 1, 1}
	k = 1
	assert.Equal(t, true, containsNearbyDuplicate(nums, k))

	nums = []int{1, 2, 3, 1, 2, 3}
	k = 2
	assert.Equal(t, false, containsNearbyDuplicate(nums, k))
}
