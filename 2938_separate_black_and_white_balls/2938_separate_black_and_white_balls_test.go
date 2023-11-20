package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMinimumSteps(t *testing.T) {
	assert.Equal(t, int64(1), minimumSteps("101"))
	assert.Equal(t, int64(2), minimumSteps("100"))
	assert.Equal(t, int64(0), minimumSteps("0111"))
	assert.Equal(t, int64(9), minimumSteps("111000"))
}
