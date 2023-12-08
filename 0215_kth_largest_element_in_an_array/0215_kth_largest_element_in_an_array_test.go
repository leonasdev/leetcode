package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindKthLargest(t *testing.T) {
	assert.Equal(t, findKthLargest([]int{3, 2, 1, 5, 6, 4}, 2), 5)
	assert.Equal(t, findKthLargest([]int{3, 2, 3, 1, 2, 4, 5, 5, 6}, 4), 4)
}
