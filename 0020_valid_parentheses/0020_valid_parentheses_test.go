package leetcode

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIsValid(t *testing.T) {
	testCases := map[string]struct {
		s        string
		expected bool
	}{
		"1": {
			s:        "()",
			expected: true,
		},
		"2": {
			s:        "()[]{}",
			expected: true,
		},
		"3": {
			s:        "(]",
			expected: false,
		},
		"4": {
			s:        "[{()}]",
			expected: true,
		},
		"5": {
			s:        "[{(})]",
			expected: false,
		},
		"6": {
			s:        "[",
			expected: false,
		},
	}
	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := isValid(tc.s)
			if actual != tc.expected {
				t.Errorf("Input: \"%v\" expected %v, got %v", tc.s, tc.expected, actual)
			}
		})
	}
}

func TestIsValidWithTestify(t *testing.T) {
	assert.Equal(t, isValid("()"), true)
	assert.Equal(t, isValid("()[]{}"), true)
	assert.Equal(t, isValid("(]"), false)
	assert.Equal(t, isValid("[{()}]"), true)
	assert.Equal(t, isValid("[{(})]"), false)
	assert.Equal(t, isValid("["), false)
}
