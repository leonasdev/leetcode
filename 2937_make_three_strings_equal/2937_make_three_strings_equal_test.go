package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFindMinimumOperations(t *testing.T) {
	assert.Equal(t, 2, findMinimumOperations("abc", "abb", "ab"))
	assert.Equal(t, -1, findMinimumOperations("dac", "bac", "cac"))
	assert.Equal(t, 3, findMinimumOperations("oxz", "oxx", "oxy"))
}
