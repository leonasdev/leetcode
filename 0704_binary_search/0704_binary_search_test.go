package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSearch(t *testing.T) {
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9, 12}, 9))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 2))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9, 12}, 111))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9}, 111))
	assert.Equal(t, -1, search([]int{-1, 0, 3, 5, 9}, 2))
	assert.Equal(t, 2, search([]int{-1, 0, 3, 5, 9}, 3))
	assert.Equal(t, 4, search([]int{-1, 0, 3, 5, 9}, 9))
	assert.Equal(t, 0, search([]int{-1, 0, 3, 5, 9}, -1))
	assert.Equal(t, 5, search([]int{-1, 0, 1, 3, 5, 9}, 9))
	assert.Equal(t, 0, search([]int{-1, 0, 1, 3, 5, 9}, -1))
}
