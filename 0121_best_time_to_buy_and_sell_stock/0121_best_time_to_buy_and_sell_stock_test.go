package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaxProfit(t *testing.T) {
	assert.Equal(t, 5, maxProfit([]int{7, 1, 5, 3, 6, 4}))

	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))

	assert.Equal(t, 2, maxProfit([]int{2, 4, 1}))

	assert.Equal(t, 2, maxProfit([]int{2, 4, 1, 2, 2, 2}))
}
