package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCharacterReplacement(t *testing.T) {
	assert.Equal(t, 4, characterReplacement("ABAB", 2))
	assert.Equal(t, 4, characterReplacement("AABABBA", 1))
	assert.Equal(t, 5, characterReplacement("AABABBA", 2))
}
