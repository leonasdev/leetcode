package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCarFleet(t *testing.T) {
	assert.Equal(t, 3, carFleet(12, []int{10, 8, 0, 5, 3}, []int{2, 4, 1, 1, 3}))
	assert.Equal(t, 1, carFleet(10, []int{3}, []int{3}))
	assert.Equal(t, 1, carFleet(100, []int{0, 2, 4}, []int{4, 2, 1}))
}
