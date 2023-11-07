package leetcode

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTwoSum(t *testing.T) {
	testCases := map[string]struct {
		nums     []int
		target   int
		expected []int
	}{
		"1": {
			nums:     []int{2, 7, 11, 15},
			target:   9,
			expected: []int{0, 1},
		},
		"2": {
			nums:     []int{3, 2, 4},
			target:   6,
			expected: []int{1, 2},
		},
		"3": {
			nums:     []int{3, 3},
			target:   6,
			expected: []int{0, 1},
		},
	}

	for name, tc := range testCases {
		t.Run(name, func(t *testing.T) {
			actual := twoSum(tc.nums, tc.target)
			if !reflect.DeepEqual(actual, tc.expected) {
				t.Fatalf("Expected %v, got %v", tc.expected, actual)
			}
		})
	}
}

func TestTwoSumWithTestify(t *testing.T) {
	assert.ElementsMatch(t, []int{1, 0}, twoSum([]int{2, 7, 11, 15}, 9))
	assert.ElementsMatch(t, []int{1, 2}, twoSum([]int{3, 2, 4}, 6))
	assert.ElementsMatch(t, []int{1, 0}, twoSum([]int{3, 3}, 6))
}
